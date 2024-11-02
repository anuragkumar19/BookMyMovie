package persons

import (
	"context"
	"strings"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
	"bookmymovie.app/bookmymovie/services/auth"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5/pgtype"
)

type SearchParams struct {
	Pagination services.PaginationParams
	Query      string
}

func (params *SearchParams) Transform() *SearchParams {
	params.Pagination = *params.Pagination.Transform()
	params.Query = strings.TrimSpace(params.Query)
	return params
}

type SearchResult struct {
	Meta    services.PaginationResult
	persons []database.SearchMoviesPersonRow
}

func (params SearchParams) Validate() error {
	return validation.ValidateStruct(
		&params,
		validation.Field(&params.Pagination),
		validation.Field(&params.Query, validation.Required, validation.Length(1, 0)),
	)
}

func (s *Persons) Search(ctx context.Context, authMeta *auth.Metadata, params *SearchParams) (SearchResult, error) {
	if err := authMeta.Valid(); err != nil {
		return SearchResult{}, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesPersonsSearch); err != nil {
		return SearchResult{}, err
	}

	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return SearchResult{}, err
		}
		return SearchResult{}, services.NewError(services.ErrorTypeInvalidArgument, err.Error())
	}

	search, err := s.db.SearchMoviesPerson(ctx, &database.SearchMoviesPersonParams{
		CreatedAt: pgtype.Timestamptz{
			Valid: true,
			Time:  params.Pagination.Timestamp,
		},
		Limit:     params.Pagination.Limit,
		Offset:    (params.Pagination.Page - 1) * params.Pagination.Limit,
		ToTsquery: params.Query,
	})
	if err != nil {
		return SearchResult{}, err
	}

	return SearchResult{
		persons: search,
		Meta: services.PaginationResult{
			Page:      params.Pagination.Page,
			Limit:     params.Pagination.Limit,
			Timestamp: params.Pagination.Timestamp,
			Count:     len(search),
		},
	}, nil
}
