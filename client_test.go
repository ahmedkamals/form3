package form3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/ahmedkamals/form3/internal/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type (
	Form3TestSuite struct {
		suite.Suite
		apiClient        *Client
		seedAccounts     []AccountData
		toCreateAccounts []AccountData
	}

	fixturesData map[string][]AccountData
)

const (
	seed     string = "seed"
	toCreate string = "toCreate"
	invalid         = "invalid"
)

func (suite *Form3TestSuite) SetupTest() {
	suite.seedAccounts = suite.loadFixtures()[seed]

	ctx := context.Background()
	for _, account := range suite.seedAccounts {
		_, err := suite.apiClient.CreateAccount(ctx, account)
		if err != nil {
			suite.Error(err)
			suite.Fail(err.Error())
		}
	}
}

func (suite *Form3TestSuite) TearDownTest() {
	ctx := context.Background()

	toDeleteAccounts := append(suite.toCreateAccounts, suite.seedAccounts...)

	for _, accountToDelete := range toDeleteAccounts {
		err := suite.apiClient.DeleteAccount(ctx, uuid.MustParse(accountToDelete.UUID), *accountToDelete.Version)
		if err != nil && !errors.Is(errors.Kind(http.StatusNotFound), err) {
			suite.Fail(err.Error())
		}
	}

	suite.toCreateAccounts = []AccountData{}
}

func (suite *Form3TestSuite) TestCreateAccount() {
	accounts := suite.loadFixtures()
	suite.toCreateAccounts = accounts[toCreate]
	invalidAccounts := accounts[invalid]

	testCases := []struct {
		id            string
		input         AccountData
		expected      *AccountData
		expectedError error
	}{
		{
			id:            "Should report that id is not a seed uuid",
			input:         invalidAccounts[0],
			expected:      nil,
			expectedError: errors.Errorf("id in body must be of type uuid: \"invalid_uuid\""),
		},
		{
			id:            "Should not able to create account - duplicate entry",
			input:         invalidAccounts[1],
			expected:      nil,
			expectedError: errors.Errorf(fmt.Sprintf("Account cannot be created as it violates a duplicate constraint")),
		},
		{
			id:            "Should be able to create the account",
			input:         suite.toCreateAccounts[0],
			expected:      &suite.toCreateAccounts[0],
			expectedError: nil,
		},
	}

	ctx := context.Background()

	for _, testCase := range testCases {
		suite.T().Run(testCase.id, func(t *testing.T) {
			accountData, err := suite.apiClient.CreateAccount(ctx, testCase.input)

			if testCase.expectedError == nil {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expected, accountData)
				return
			}

			assert.NotNil(t, err, err.Error())
			assert.Contains(t, err.Error(), testCase.expectedError.Error())
		})
	}
}

func (suite *Form3TestSuite) TestFetchAccount() {
	notExistAccountUUID, _ := uuid.NewUUID()
	testCases := []struct {
		id            string
		input         uuid.UUID
		expected      *AccountData
		expectedError error
	}{
		{
			id:            "Should not able to fetch account",
			input:         notExistAccountUUID,
			expected:      nil,
			expectedError: errors.Errorf(fmt.Sprintf("record %s does not exist", notExistAccountUUID)),
		},
		{
			id:            "Should be able to fetch the account",
			input:         uuid.MustParse(suite.seedAccounts[0].UUID),
			expected:      &suite.seedAccounts[0],
			expectedError: nil,
		},
	}

	ctx := context.Background()

	for _, testCase := range testCases {
		suite.T().Run(testCase.id, func(t *testing.T) {
			accountData, err := suite.apiClient.FetchAccount(ctx, testCase.input)

			if testCase.expectedError == nil {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expected, accountData)
				return
			}

			assert.NotNil(t, err, err.Error())
			assert.Equal(t, testCase.expectedError.Error(), err.Error())
		})
	}
}

func (suite *Form3TestSuite) TestDeleteAccount() {
	notExistAccountUUID, _ := uuid.NewUUID()
	testCases := []struct {
		id            string
		input         map[uuid.UUID]uint64
		expectedError error
	}{
		{
			id:            "Should not able to delete account - UUID does not exist",
			input:         map[uuid.UUID]uint64{notExistAccountUUID: 0},
			expectedError: errors.Errorf("EOF"),
		},
		{
			id:            "Should not able to delete account - version does not exist",
			input:         map[uuid.UUID]uint64{uuid.MustParse(suite.seedAccounts[0].UUID): 1},
			expectedError: errors.Errorf("invalid version"),
		},
		{
			id:            "Should be able to find the account",
			input:         map[uuid.UUID]uint64{uuid.MustParse(suite.seedAccounts[0].UUID): 0},
			expectedError: nil,
		},
	}

	ctx := context.Background()

	for _, testCase := range testCases {
		suite.T().Run(testCase.id, func(t *testing.T) {
			for uuidValue, version := range testCase.input {
				err := suite.apiClient.DeleteAccount(ctx, uuidValue, version)

				if testCase.expectedError == nil {
					assert.Nil(t, err)
					return
				}

				assert.NotNil(t, err, err.Error())
				assert.Equal(t, testCase.expectedError.Error(), err.Error())
			}
		})
	}
}

func (suite *Form3TestSuite) loadFixtures() fixturesData {
	config := Config{
		endpoint: os.Getenv("API_ENDPOINT"),
	}

	suite.apiClient = NewClient(config, &http.Client{})

	fixtures, err := os.Open(os.Getenv("FIXTURES_PATH"))
	if err != nil {
		suite.Fail(err.Error())
	}

	var fixturesData fixturesData

	err = json.NewDecoder(fixtures).Decode(&fixturesData)
	if err != nil {
		suite.Fail(err.Error())
	}

	return fixturesData
}

func TestForm3TestSuite(t *testing.T) {
	suite.Run(t, new(Form3TestSuite))
}
