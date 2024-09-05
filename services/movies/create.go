package movies

import (
	"context"

	"bookmymovie.app/bookmymovie/services/auth"
)

type CreateMovieParams struct {
}

func (s *Movies) CreateMovie(_ context.Context, authMeta *auth.Metadata, _ *CreateMovieParams) (id int64, err error) {
	if err := authMeta.Valid(); err != nil {
		return 0, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MovieCreate); err != nil {
		return 0, err
	}
	return 1, nil
}
