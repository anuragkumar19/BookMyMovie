package services

import "time"

type ErrorType string

const (
	ErrorTypeInvalidArgument   ErrorType = "invalid_argument"
	ErrorTypeNotFound          ErrorType = "not_found"
	ErrorTypeAlreadyExist      ErrorType = "already_exist"
	ErrorTypePermissionDenied  ErrorType = "permission_denied"
	ErrorTypeResourceExhausted ErrorType = "resource_exhausted"
	ErrorTypeConflict          ErrorType = "conflict"
	ErrorTypeUnauthenticated   ErrorType = "unauthenticated"
)

type Error struct {
	errorType ErrorType
	msg       string
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Type() ErrorType {
	return e.errorType
}

func NewError(errorType ErrorType, msg string) *Error {
	return &Error{
		errorType: errorType,
		msg:       msg,
	}
}

func NewRateLimitErrorMessage(tryAfter time.Duration, message string) string {
	msg := "too many requests"
	if message != "" {
		msg = message
	}
	return msg + ", try after " + tryAfter.String()
}
