package storage

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type StorageConfig struct {
	FirebaseAdminConfigJSON []byte
	Bucket                  string
}

func (config *StorageConfig) Validate() error {
	return validation.ValidateStruct(
		config,
		validation.Field(&config.FirebaseAdminConfigJSON, validation.Required),
		validation.Field(&config.Bucket, validation.Required),
	)
}
