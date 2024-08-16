package storage

import (
	"context"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog"
)

type Storage struct {
	logger *zerolog.Logger
	config *StorageConfig
	client *minio.Client
}

func New(config *StorageConfig, logger *zerolog.Logger) Storage {
	if err := config.Validate(); err != nil {
		logger.Fatal().Err(err).Msg("storage config validation failed")
	}

	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.Secret, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create storage client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	exist, err := client.BucketExists(ctx, config.Bucket)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to check check bucket existence")
	}
	if !exist {
		if !config.AutoCreateBucket {
			logger.Fatal().Err(err).Bool("auto_create_bucket", config.AutoCreateBucket).Str("bucket", config.Bucket).Msg("bucket with specified name doesn't exits")
		}
		logger.Info().Bool("auto_create_bucket", config.AutoCreateBucket).Str("bucket", config.Bucket).Msg("bucket doesn't exist creating bucket")
		if err := client.MakeBucket(ctx, config.Bucket, minio.MakeBucketOptions{}); err != nil {
			logger.Fatal().Err(err).Bool("auto_create_bucket", config.AutoCreateBucket).Str("bucket", config.Bucket).Msg("failed to create bucket")
		}
		logger.Info().Bool("auto_create_bucket", config.AutoCreateBucket).Str("bucket", config.Bucket).Msg("bucket created")
	}

	return Storage{
		config: config,
		logger: logger,
		client: client,
	}
}
