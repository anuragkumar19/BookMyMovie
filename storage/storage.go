package storage

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/rs/zerolog"
	"google.golang.org/api/option"
)

type Storage struct {
	logger *zerolog.Logger
	config *StorageConfig
	bucket *storage.BucketHandle
}

func New(config *StorageConfig, logger *zerolog.Logger) Storage {
	if err := config.Validate(); err != nil {
		logger.Fatal().Err(err).Msg("storage config validation failed")
	}

	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_ADMIN_CONFIG")))
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		StorageBucket: os.Getenv("STORAGE_BUCKET"),
	}, opt)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create firebase app")
	}

	s, err := app.Storage(context.Background())
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to init firebase app storage")
	}
	buckerHandle, err := s.DefaultBucket()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to init default bucked handle")
	}

	return Storage{
		config: config,
		logger: logger,
		bucket: buckerHandle,
	}
}
