package serviceserrors

type ErrorType string

const (
	ErrorTypeInvalidArgument  ErrorType = "invalid_argument"
	ErrorTypeNotFound         ErrorType = "not_found"
	ErrorTypeAlreadyExist     ErrorType = "already_exist"
	ErrorTypePermissionDenied ErrorType = "permission_denied"
	ErrorResourceExhausted    ErrorType = "resource_exhausted"
	ErrorConflict             ErrorType = "conflict"
	ErrorUnauthenticated      ErrorType = "unauthenticated"
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

func New(errorType ErrorType, msg string) *Error {
	return &Error{
		errorType: errorType,
		msg:       msg,
	}
}
