package errors

import (
	"errors"
)

var (
	ErrUpdateConflict = errors.New("update conflict - version mismatch")
	ErrOTPExpired     = errors.New("otp expired")
	ErrOTPMismatch    = errors.New("otp mismatch")
)
