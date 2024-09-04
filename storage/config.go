package storage

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Config struct {
	Endpoint         string
	AccessKey        string
	Secret           string
	UseSSL           bool
	Bucket           string
	AutoCreateBucket bool
	Region           string
}

func (config *Config) Validate() error {
	return validation.ValidateStruct(
		config,
		validation.Field(&config.Endpoint, validation.Required),
		validation.Field(&config.AccessKey, validation.Required),
		validation.Field(&config.Secret, validation.Required),
		validation.Field(&config.UseSSL),
		validation.Field(&config.Bucket, validation.Required),
		validation.Field(&config.Region),
		validation.Field(&config.AutoCreateBucket),
	)
}
