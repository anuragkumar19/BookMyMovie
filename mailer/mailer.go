package mailer

import (
	"sync"

	"github.com/rs/zerolog"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	logger *zerolog.Logger
	dialer *gomail.Dialer
	config *Config
	wg     sync.WaitGroup
}

func New(config *Config, logger *zerolog.Logger) Mailer {
	if err := config.Validate(); err != nil {
		logger.Fatal().Err(err).Msg("mailer config validation failed")
	}

	dialer := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	if _, err := dialer.Dial(); err != nil {
		logger.Fatal().Err(err).Msg("authentication with smtp server failed")
	}

	return Mailer{
		dialer: dialer,
		config: config,
		wg:     sync.WaitGroup{},
		logger: logger,
	}
}

func (m *Mailer) Wait() {
	m.wg.Wait()
}
