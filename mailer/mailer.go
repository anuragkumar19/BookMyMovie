package mailer

import (
	"errors"
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

func New(config *Config, logger *zerolog.Logger) (Mailer, error) {
	if err := config.Validate(); err != nil {
		return Mailer{}, errors.Join(errors.New("mailer config validation failed"), err)
	}

	dialer := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	if _, err := dialer.Dial(); err != nil {
		return Mailer{}, errors.Join(errors.New("authentication with smtp server failed"), err)
	}

	return Mailer{
		dialer: dialer,
		config: config,
		wg:     sync.WaitGroup{},
		logger: logger,
	}, nil
}

func (m *Mailer) Wait() {
	m.wg.Wait()
}
