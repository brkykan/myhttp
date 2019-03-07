package errors

import "errors"

const (
	connectionTimedOutError = "connection timed out"
	invalidURL              = "invalid URL"
	connectionFailed        = "connection failed"
)

var ErrConnectionTimedOut = errors.New(connectionTimedOutError)
var ErrInvalidURL = errors.New(invalidURL)
var ErrConnectionFailed = errors.New(connectionFailed)
