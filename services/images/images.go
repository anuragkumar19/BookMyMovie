package images

import (
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/storage"
	"github.com/rs/zerolog"
)

type Images struct {
	config  *Config
	logger  *zerolog.Logger
	storage *storage.Storage
	auth    *auth.Auth
}

func New(config *Config, logger *zerolog.Logger, a *auth.Auth, s *storage.Storage) Images {
	return Images{
		config:  config,
		logger:  logger,
		storage: s,
		auth:    a,
	}
}
