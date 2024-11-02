package images

import (
	"context"

	"bookmymovie.app/bookmymovie/services"
	"bookmymovie.app/bookmymovie/services/auth"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ListResult struct {
	Meta   services.PaginationResult
	Images []ImageDetail
}

type ListParams struct {
	Pagination services.PaginationParams
}

func (params *ListParams) Transform() *ListParams {
	params.Pagination = *params.Pagination.Transform()
	return params
}

func (params ListParams) Validate() error {
	return validation.ValidateStruct(
		&params,
		validation.Field(&params.Pagination),
	)
}

func (s *Images) List(ctx context.Context, authMeta *auth.Metadata, params ListParams) (ListResult, error) {
	if err := authMeta.Valid(); err != nil {
		return ListResult{}, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.ImagesList); err != nil {
		return ListResult{}, err
	}

	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return ListResult{}, err
		}
		return ListResult{}, services.NewError(services.ErrorTypeInvalidArgument, err.Error())
	}

	return ListResult{}, nil
}
