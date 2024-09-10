package auth

import (
	"fmt"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
)

type Permission string

type rolePermissionMap map[Permission]bool

const (
	MovieCreate Permission = "movie:create"
	MovieUpdate Permission = "movie:update"
	MovieDelete Permission = "movie:delete"

	MoviesLanguagesCreate Permission = "movies:languages:create"
	MoviesLanguagesDelete Permission = "movies:languages:delete"
	MoviesLanguagesUpdate Permission = "movies:languages:update"

	MoviesGenresCreate Permission = "movies:genres:create"
	MoviesGenresDelete Permission = "movies:genres:delete"
	MoviesGenresUpdate Permission = "movies:genres:update"

	MoviesFormatsCreate Permission = "movies:formats:create"
	MoviesFormatsDelete Permission = "movies:formats:delete"
	MoviesFormatsUpdate Permission = "movies:formats:update"

	MoviesPersonsCreate Permission = "movies:persons:create"
	MoviesPersonsUpdate Permission = "movies:persons:update"
	MoviesPersonsDelete Permission = "movies:persons:delete"
)

var adminPermission = rolePermissionMap{
	MovieCreate:           true,
	MovieUpdate:           true,
	MovieDelete:           true,
	MoviesLanguagesCreate: true,
	MoviesLanguagesDelete: true,
	MoviesLanguagesUpdate: true,
	MoviesGenresCreate:    true,
	MoviesGenresDelete:    true,
	MoviesGenresUpdate:    true,
	MoviesPersonsCreate:   true,
	MoviesPersonsUpdate:   true,
	MoviesPersonsDelete:   true,
	MoviesFormatsCreate:   true,
	MoviesFormatsDelete:   true,
	MoviesFormatsUpdate:   true,
}

var regularUserPermission = rolePermissionMap{}

func (*Auth) CheckPermissions(authMeta *Metadata, ps ...Permission) error {
	if err := authMeta.Valid(); err != nil {
		return err
	}
	switch authMeta.UserRole() {
	case database.RolesAdmin:
		return checkPermission(adminPermission, ps)
	case database.RolesRegularUser:
		return checkPermission(regularUserPermission, ps)
	default:
		panic("invalid role")
	}
}

func checkPermission(m rolePermissionMap, ps []Permission) error {
	for _, p := range ps {
		if !m[p] {
			return serviceserrors.New(serviceserrors.ErrorTypePermissionDenied, fmt.Sprintf("user doesn't have permission %s", p))
		}
	}

	return nil
}
