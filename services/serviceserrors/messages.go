package serviceserrors

import (
	"time"
)

func NewRateLimitMessage(tryAfter time.Duration, message string) string {
	msg := "too many requests"
	if message != "" {
		msg = message
	}
	return msg + ", try after " + tryAfter.String()
}
