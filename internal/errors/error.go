package errors

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

type (
	// Kind is the error classification.
	Kind int
	// Operation that was performed to produce this error.
	Operation string
	// Error is the implementation of the error interface.
	Error struct {
		// Kind of the error such as entity not found
		// or "Other" if its class is unknown or irrelevant.
		Kind Kind
		// Op is the performed operation, usually it is the method name.
		// e.g. Bus.Publish()
		Op Operation
		// Err is the underlying error that has triggered this one, if any.
		Err error
	}
)

const (
	// Other is a fallback classification of the error.
	Other Kind = 0
	// Invalid input is not allowed.
	Invalid Kind = 1 << iota
	// NotFound item is not found.
	NotFound
	// MinLength should be 1 or more.
	MinLength
	// Failure to apply an operation.
	Failure
	// Panic recovery errors.
	Panic
)

var (
	// Separator is used to separate nested errors.
	Separator = ":\n\t"
	kinds     = map[Kind]string{
		Other:     "NA",
		Invalid:   "invalid input is provided",
		NotFound:  "not found",
		MinLength: "not matching minimum length constraint",
		Failure:   "could not perform operation on the provided data %s",
		Panic:     "panic",
	}
)

// String will try locally defined kinds, and if it did not find anything,
// it would fall back to http error codes.
func (k Kind) String() string {
	if kind, ok := kinds[k]; ok {
		return kind
	}

	return http.StatusText(int(k))
}

// E creates an error
func E(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to errors.E with no arguments")
	}

	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Operation:
			e.Op = arg
		case Kind:
			e.Kind = arg
		case *Error:
			clone := *arg
			e.Err = &clone
		case error:
			e.Err = arg
		default:
			_, file, line, _ := runtime.Caller(1)
			log.Printf("errors.E: bad call from %s:%d: %v", file, line, args)

			return Errorf("unknown type %T, value %v in error call", arg, arg)
		}
	}

	return e
}

// FullError returns the full error string representation.
func (e *Error) FullError() string {
	b := new(bytes.Buffer)

	if e.Op != "" {
		pad(b, ": ")
		b.WriteString(string(e.Op))
	}

	if e.Kind != 0 {
		pad(b, ": ")
		b.WriteString(e.Kind.String())
	}

	e.errors(b, Separator)

	if b.Len() == 0 {
		return "no error"
	}

	return b.String()
}

func (e *Error) Error() string {
	b := new(bytes.Buffer)
	e.errors(b, Separator)

	return b.String()
}

func (e *Error) errors(b *bytes.Buffer, separator string) {
	if e.Err == nil {
		return
	}

	prevErr, ok := e.Err.(*Error)
	if !ok {
		pad(b, ": ")
		b.WriteString(e.Err.Error())

		return
	}

	// Indent on new line if we are cascading non-empty errors.
	if !prevErr.isZero() {
		pad(b, separator)
		b.WriteString(e.Err.Error())
	}
}

func (e *Error) isZero() bool {
	return e.Op == "" && e.Kind == 0 && e.Err == nil
}

// Errorf creates an error from a given format string.
func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

// Is checks if the error of a given kind.
func Is(kind Kind, err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}

	if e.Kind != Other {
		return e.Kind == kind
	}

	if e.Err != nil {
		return Is(kind, e.Err)
	}

	return false
}

// pad appends str to the buffer if the buffer already has some data.
func pad(b *bytes.Buffer, str string) {
	if b.Len() == 0 {
		return
	}

	b.WriteString(str)
}
