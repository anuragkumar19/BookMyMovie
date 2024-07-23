package auth

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type AuthConfig struct {
	Host string

	AccessTokenSecret   string
	RefreshTokenSecret  string
	AccessTokenLifetime time.Duration

	LoginOTPSendingRate           int
	LoginOTPSendingRateTimeWindow time.Duration
	LoginOTPLifetime              time.Duration
	MaxOTPIncorrectAttempts       int
}

func (config *AuthConfig) Validate() error {
	return validation.ValidateStruct(
		config,
		validation.Field(&config.Host, validation.Required, is.URL),
		validation.Field(&config.AccessTokenSecret, validation.Required),
		validation.Field(&config.RefreshTokenSecret, validation.Required),
		validation.Field(&config.AccessTokenLifetime, validation.Required),
		validation.Field(&config.LoginOTPSendingRate, validation.Required, validation.Min(1)),
		validation.Field(&config.LoginOTPSendingRateTimeWindow, validation.Required, validation.Min(time.Millisecond)),
		validation.Field(&config.LoginOTPLifetime, validation.Required, validation.Min(time.Millisecond)),
		validation.Field(&config.MaxOTPIncorrectAttempts, validation.Required, validation.Min(1)),
	)
}

func DefaultConfig() AuthConfig {
	return AuthConfig{
		Host:                          "",
		AccessTokenSecret:             "",
		RefreshTokenSecret:            "",
		AccessTokenLifetime:           time.Minute,
		LoginOTPSendingRate:           10,
		LoginOTPSendingRateTimeWindow: time.Hour,
		LoginOTPLifetime:              10 * time.Minute,
		MaxOTPIncorrectAttempts:       10,
	}
}
