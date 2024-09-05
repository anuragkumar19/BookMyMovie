package mailer

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string

	FromAddress     string
	FromDisplayName string
	ReplyTo         string
}

func (config *Config) Validate() error {
	return validation.ValidateStruct(
		config,
		validation.Field(&config.Host, validation.Required, is.Host),
		validation.Field(&config.Port, validation.Required), // is.Port,

		validation.Field(&config.Username, validation.Required),
		validation.Field(&config.Password, validation.Required),
		validation.Field(&config.FromAddress, validation.Required, is.Email),
		validation.Field(&config.FromDisplayName),
		validation.Field(&config.ReplyTo, is.Email),
	)
}
