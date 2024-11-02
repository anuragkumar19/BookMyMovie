package storage

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (s *Storage) ListObjects(ctx context.Context, limit int) {
	s.Client.ListObjects(ctx, s.config.Bucket, minio.ListObjectsOptions{
		Recursive: true,
	})
}
