package form3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ahmedkamals/form3/internal/errors"
	"github.com/google/uuid"
)

type (
	// Client of the api.
	Client struct {
		config     Config
		httpClient *http.Client
	}
)

const (
	defaultTimeout = 10 * time.Second
)

// NewClient creates an api Client instance.
func NewClient(config Config, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: defaultTimeout,
		}
	}

	return &Client{
		config:     config,
		httpClient: httpClient,
	}
}

// CreateAccount calls create account endpoint.
func (c *Client) CreateAccount(ctx context.Context, accountData AccountData) (*AccountData, error) {
	const op errors.Operation = "Client.FetchAccount"

	body := struct {
		Data AccountData `json:"data"`
	}{
		Data: accountData,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.E(op, err)
	}

	reqBodyBuffer := bytes.NewBuffer(jsonBody)

	req, err := c.buildRequest(
		http.MethodPost,
		fmt.Sprintf("%s/v1/organisation/accounts", c.config.endpoint),
		reqBodyBuffer,
	)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return c.Do(ctx, req)
}

// FetchAccount calls fetch account endpoint.
func (c *Client) FetchAccount(ctx context.Context, accountUUID uuid.UUID) (*AccountData, error) {
	const op errors.Operation = "Client.FetchAccount"

	req, err := c.buildRequest(
		http.MethodGet,
		fmt.Sprintf("%s/v1/organisation/accounts/%s", c.config.endpoint, accountUUID),
		nil,
	)

	if err != nil {
		return nil, errors.E(op, err)
	}

	return c.Do(ctx, req)
}

// DeleteAccount calls delete account endpoint.
func (c *Client) DeleteAccount(ctx context.Context, accountUUID uuid.UUID, version uint64) error {
	const op errors.Operation = "Client.DeleteAccount"

	req, err := c.buildRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/v1/organisation/accounts/%s?version=%d", c.config.endpoint, accountUUID, version),
		nil,
	)

	if err != nil {
		return errors.E(op, err)
	}

	_, err = c.Do(ctx, req)
	if err != nil {
		return errors.E(op, err)
	}

	return nil
}

// Do performs http requests.
func (c *Client) Do(ctx context.Context, req *http.Request) (*AccountData, error) {
	const op errors.Operation = "Client.Do"

	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.E(op, err)
	}

	defer resp.Body.Close()

	err = c.scanForErrors(resp)
	if err != nil {
		return nil, errors.E(op, errors.Kind(resp.StatusCode), err)
	}

	if req.Method == http.MethodDelete {
		return nil, nil
	}

	var response struct {
		Data *AccountData `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return response.Data, nil
}

func (c *Client) buildRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Add("Date", time.Now().Format(time.RFC1123))
	req.Header.Add("Host", c.config.endpoint)

	return req, nil
}

func (c *Client) scanForErrors(resp *http.Response) error {
	if resp.StatusCode < 300 {
		return nil
	}

	const op errors.Operation = "Client.scanForErrors"

	var data struct {
		ErrorMessage string `json:"error_message"`
	}
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return errors.E(op, err)
	}

	return errors.E(op, errors.Errorf(data.ErrorMessage))
}
