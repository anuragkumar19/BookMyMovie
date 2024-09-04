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

	MoviesGenresCreate Permission = "movies:genres:create"
	MoviesGenresDelete Permission = "movies:genres:delete"

	MoviesFormatsCreate Permission = "movies:formats:create"
	MoviesFormatsDelete Permission = "movies:formats:delete"

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
	MoviesGenresCreate:    true,
	MoviesGenresDelete:    true,
	MoviesPersonsCreate:   true,
	MoviesPersonsUpdate:   true,
	MoviesPersonsDelete:   true,
	MoviesFormatsCreate:   true,
	MoviesFormatsDelete:   true,
}

var regularUserPermission = rolePermissionMap{}

func (*Auth) CheckPermissions(auth *Metadata, ps ...Permission) error {
	switch auth.UserRole {
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
			return serviceserrors.UnauthorizedError(fmt.Errorf("user doesn't have permission %s", p))
		}
	}

	return nil
}
