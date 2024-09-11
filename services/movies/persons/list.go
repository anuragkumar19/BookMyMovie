package persons

import (
	"context"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
	"bookmymovie.app/bookmymovie/services/auth"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5/pgtype"
)

type ListParams struct {
	Pagination services.PaginationParams
}

func (params *ListParams) Transform() *ListParams {
	params.Pagination = *params.Pagination.Transform()
	return params
}

type ListResult struct {
	Meta    services.PaginationResult
	persons []database.ListMoviesPersonRow
}

func (params ListParams) Validate() error {
	return validation.ValidateStruct(
		&params,
		validation.Field(&params.Pagination),
	)
}

func (s *Persons) List(ctx context.Context, authMeta *auth.Metadata, params *ListParams) (ListResult, error) {
	if err := authMeta.Valid(); err != nil {
		return ListResult{}, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesPersonsList); err != nil {
		return ListResult{}, err
	}

	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return ListResult{}, err
		}
		return ListResult{}, services.NewError(services.ErrorTypeInvalidArgument, err.Error())
	}

	list, err := s.db.ListMoviesPerson(ctx, &database.ListMoviesPersonParams{
		CreatedAt: pgtype.Timestamptz{
			Valid: true,
			Time:  params.Pagination.Timestamp,
		},
		Limit:  params.Pagination.Limit,
		Offset: (params.Pagination.Page - 1) * params.Pagination.Limit,
	})
	if err != nil {
		return ListResult{}, err
	}

	return ListResult{
		persons: list,
		Meta: services.PaginationResult{
			Page:      params.Pagination.Page,
			Limit:     params.Pagination.Limit,
			Timestamp: params.Pagination.Timestamp,
			Count:     len(list),
		},
	}, nil
}
