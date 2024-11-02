package storage

import (
	"context"
	"errors"

	"github.com/minio/minio-go/v7"
)

func (s *Storage) Exist(ctx context.Context, key string) (bool, error) {
	_, err := s.Client.StatObject(ctx, s.config.Bucket, key, minio.StatObjectOptions{})

	if err != nil {
		if errors.As(err, &minio.ErrorResponse{}) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
