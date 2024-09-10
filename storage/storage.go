package storage

import (
	"context"
	"errors"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog"
)

type Storage struct {
	logger *zerolog.Logger
	config *Config
	client *minio.Client
}

func New(ctx context.Context, config *Config, logger *zerolog.Logger) (Storage, error) {
	if err := config.Validate(); err != nil {
		return Storage{}, errors.Join(errors.New("storage config validation failed"), err)
	}

	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.Secret, ""),
		Secure: config.UseSSL,
		Region: config.Region,
	})
	if err != nil {
		return Storage{}, errors.Join(errors.New("failed to create storage client"), err)
	}

	exist, err := client.BucketExists(ctx, config.Bucket)
	if err != nil {
		return Storage{}, errors.Join(errors.New("failed to check check bucket existence"), err)
	}
	if !exist {
		if !config.AutoCreateBucket {
			return Storage{}, errors.Join(errors.New("bucket with specified name doesn't exits"), err)
		}
		logger.Info().Bool("auto_create_bucket", config.AutoCreateBucket).Str("bucket", config.Bucket).Msg("bucket doesn't exist creating bucket")
		if err := client.MakeBucket(ctx, config.Bucket, minio.MakeBucketOptions{}); err != nil {
			return Storage{}, errors.Join(errors.New("failed to create bucket"), err)
		}
		logger.Info().Bool("auto_create_bucket", config.AutoCreateBucket).Str("bucket", config.Bucket).Msg("bucket created")
	}

	return Storage{
		config: config,
		logger: logger,
		client: client,
	}, nil
}
