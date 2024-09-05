package mailer

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	Host            string `conf:"env:HOST,flag:host"`
	Port            int    `conf:"env:PORT,flag:port"`
	Username        string `conf:"env:USERNAME,flag:username"`
	Password        string `conf:"env:PASSWORD,flag:password"`
	FromAddress     string `conf:"env:FROM_ADDRESS,flag:from-address"`
	FromDisplayName string `conf:"env:FROM_DISPLAY_NAME,flag:display-name"`
	ReplyTo         string `conf:"env:REPLY_TO,flag:reply-to"`
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
