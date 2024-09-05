package auth

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	AppPublicHost string

	AccessTokenSecret    string
	AccessTokenLifetime  time.Duration
	RefreshTokenLifetime time.Duration

	LoginOTPSendingRate           int
	LoginOTPSendingRateTimeWindow time.Duration
	LoginOTPLifetime              time.Duration
	MaxOTPIncorrectAttempts       int
}

func (config *Config) Validate() error {
	return validation.ValidateStruct(
		config,
		validation.Field(&config.AppPublicHost, validation.Required, is.URL),
		validation.Field(&config.AccessTokenSecret, validation.Required),
		validation.Field(&config.AccessTokenLifetime, validation.Required),
		validation.Field(&config.RefreshTokenLifetime, validation.Required),
		validation.Field(&config.LoginOTPSendingRate, validation.Required, validation.Min(1)),
		validation.Field(&config.LoginOTPSendingRateTimeWindow, validation.Required, validation.Min(time.Millisecond)),
		validation.Field(&config.LoginOTPLifetime, validation.Required, validation.Min(time.Millisecond)),
		validation.Field(&config.MaxOTPIncorrectAttempts, validation.Required, validation.Min(1)),
	)
}

func DefaultConfig() Config {
	return Config{
		AppPublicHost:                 "",
		AccessTokenSecret:             "",
		AccessTokenLifetime:           time.Minute,
		LoginOTPSendingRate:           10,
		LoginOTPSendingRateTimeWindow: time.Hour,
		RefreshTokenLifetime:          30 * 24 * time.Hour,
		LoginOTPLifetime:              10 * time.Minute,
		MaxOTPIncorrectAttempts:       10,
	}
}
