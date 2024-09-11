package services

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PaginationParams struct {
	Page      int32
	Limit     int32
	Timestamp time.Time
}

func (params PaginationParams) Validate() error {
	return validation.ValidateStruct(
		&params,
		validation.Field(&params.Page, validation.Required, validation.Min(1)),
		validation.Field(&params.Limit, validation.Required, validation.Min(1), validation.Max(100)),
		validation.Field(&params.Timestamp, validation.Max(time.Now())),
	)
}

func (params *PaginationParams) Transform() *PaginationParams {
	if params.Timestamp.IsZero() {
		params.Timestamp = time.Now()
	}
	return params
}

type PaginationResult struct {
	Page      int32
	Limit     int32
	Timestamp time.Time
	Count     int
}
