package errors

import (
	"errors"
	"fmt"
)

// newEmptyCredentialsError creates error for empty credentials
func newEmptyCredentialsError() error {
	return errors.New("Login credentials should not be empty")
}

// newWrongResponseStatusError creates error for invalid status of http response
func newWrongResponseStatusError(status string) error {
	return fmt.Errorf("Wrong response status: \"%s\"\n", status)
}

// newResponseError creates error for response error
func newResponseError(errorMsg string) error {
	return fmt.Errorf("Something went wrong: \"%s\"\n", errorMsg)
}
