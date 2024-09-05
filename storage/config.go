package storage

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Config struct {
	Endpoint         string `conf:"env:ENDPOINT,flag:endpoint"`
	AccessKey        string `conf:"env:ACCESS_KEY,flag:access-key"`
	Secret           string `conf:"env:SECRET,flag:secret"`
	UseSSL           bool   `conf:"env:USE_SSL,flag:use-ssl,default:false"`
	Bucket           string `conf:"env:BUCKET,flag:bucket"`
	AutoCreateBucket bool   `conf:"env:AUTO_CREATE_BUCKET,flag:auto-create-bucket,default:false"`
	Region           string `conf:"evn:REGION,flag:region"`
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
