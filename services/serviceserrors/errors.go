package serviceserrors

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	ErrAlreadyExist     = errors.New("already exist")
	ErrNotFound         = errors.New("not found")
	ErrValidationFailed = errors.New("validation error")
	ErrUpdateConflict   = errors.New("update conflict - version mismatch")
	ErrOTPExpired       = errors.New("otp expired")
	ErrOTPMismatch      = errors.New("otp mismatch")
	ErrUnauthorized     = errors.New("unauthorized")
)

func ValidationError(err validation.Errors) error {
	return errors.Join(ErrValidationFailed, err)
}

func UnauthorizedError(err error) error {
	return errors.Join(ErrUnauthorized, err)
}
