package storage

import (
	"context"
	"errors"
	"io"

	"github.com/minio/minio-go/v7"
)

func (s *Storage) Upload(ctx context.Context, src io.Reader, extension string, size int64) (key string, err error) {
	objKey, err := randObjKey()
	if err != nil {
		return "", err
	}
	info, err := s.client.PutObject(ctx, s.config.Bucket, objKey+extension, src, size, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return info.Key, nil
}

func (s *Storage) Exist(ctx context.Context, key string) (bool, error) {
	_, err := s.client.StatObject(ctx, s.config.Bucket, key, minio.StatObjectOptions{})

	if err != nil {
		if errors.As(err, &minio.ErrorResponse{}) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
