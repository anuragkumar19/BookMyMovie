package auth

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	AppPublicHost string `conf:"env:APP_PUBLIC_HOST,flag:app-public-host"`

	AccessTokenSecret    string        `conf:"env:ACCESS_TOKEN_SECRET,flag:access-token-secret"`
	AccessTokenLifetime  time.Duration `conf:"env:ACCESS_TOKEN_LIFETIME,flag:access-token-lifetime,default:1m"`
	RefreshTokenLifetime time.Duration `conf:"env:REFRESH_TOKEN_LIFETIME,flag:refresh-token-lifetime,default:30d"`

	LoginOTPSendingRate           int           `conf:"env:LOGIN_OTP_SENDING_RATE,flag:login-opt-sending-rate,default:10"`
	LoginOTPSendingRateTimeWindow time.Duration `conf:"env:LOGIN_OTP_SENDING_RATE_TIME_WINDOW,flag:login-otp-sending-rate-time-window,default:1h"`
	LoginOTPLifetime              time.Duration `conf:"env:LOGIN_OTP_LIFETIME,flag:login-otp-lifetime,default:10m"`
	MaxOTPIncorrectAttempts       int           `conf:"env:MAX_OTP_INCORRECT_ATTEMPTS,flag:max-otp-incorrect-attempts,default:10"`
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
