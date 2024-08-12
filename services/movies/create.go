package movies

import (
	"context"

	"bookmymovie.app/bookmymovie/services/auth"
)

type CreateMovieParams struct {
	AccessToken string
}

func (s *Movies) CreateMovie(ctx context.Context, params *CreateMovieParams) (id int64, err error) {
	user, err := s.auth.GetAuthMetadata(params.AccessToken)
	if err != nil {
		return 0, err
	}
	if err := s.auth.CheckPermissions(&user, auth.MovieCreate); err != nil {
		return 0, err
	}
	return 1, nil
}
