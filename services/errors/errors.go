package errors

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	ErrValidationFailed = errors.New("validation error")
	ErrUpdateConflict   = errors.New("update conflict - version mismatch")
	ErrOTPExpired       = errors.New("otp expired")
	ErrOTPMismatch      = errors.New("otp mismatch")
)

func ValidationError(err validation.Errors) error {
	return errors.Join(ErrValidationFailed, err)
}
