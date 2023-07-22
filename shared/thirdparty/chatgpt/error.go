package chatgpt

import (
	"github.com/pkg/errors"
)

var (
	// ErrorServiceUnavailable is the error when OpenAI cannot process the request due to overloaded
	ErrorServiceUnavailable = errors.New("OpenAI service is currently unavailable")
	// ErrorInvalidParameter is the error when API call returns parameter validation error.
	ErrorInvalidParameter = errors.New("OpenAI invalid parameter")
	// ErrorInvalidResponseStructure is the error when response decoding mismatch happened
	ErrorInvalidResponseStructure = errors.New("An invalid response structure was provided")
	// ErrorAPIKeyInvalid is the error when OpenAI API returns unauthorized due to invalid API key
	ErrorAPIKeyInvalid = errors.New("Invalid API Key, unauthorized")
	// ErrorUnknown is the error when other error happened
	ErrorUnknown = errors.New("Unknown error occurred")
	// ErrorModelNotFound is the error when client give unknown model
	ErrorModelNotFound = errors.New("Model not found")
	// ErrorAnswerCountLimit is the error when client give answer count more than limit
	ErrorAnswerCountLimit = errors.New("Answer count more than limit")
	// ErrorClientTimeout is the error when client took too long for the OpenAI server to respond
	ErrorClientTimeout = errors.New("Error client timeout")
	// ErrorTooManyRequest is the error when our service sent too many request to OpenAI server
	ErrorTooManyRequest = errors.New("Error too many request")
)
