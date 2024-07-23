package errors

import (
	"fmt"
	"time"
)

type RateLimitError struct {
	TryAfter                time.Duration `json:"try_after"`
	TotalRequests           int           `json:"total_requests"`
	LastSuccessfulRequestAt time.Time     `json:"last_successful_request_at"`
}

func (e RateLimitError) Error() string {
	return fmt.Sprintf("too many requests : rate limited : try after %s", e.TryAfter.String())
}

func NewRateLimitError(tryAfter time.Duration, totalRequest int, lastReqAt time.Time) *RateLimitError {
	return &RateLimitError{
		TryAfter:                tryAfter,
		TotalRequests:           totalRequest,
		LastSuccessfulRequestAt: lastReqAt,
	}
}
