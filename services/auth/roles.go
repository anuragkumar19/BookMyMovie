package auth

import (
	"fmt"

	"bookmymovie.app/bookmymovie/database"
	services_errors "bookmymovie.app/bookmymovie/services/errors"
)

type Permission string

type rolePermissionMap map[Permission]bool

const (
	MovieCreate Permission = "movie:create"
	MovieUpdate Permission = "movie:update"
	MovieDelete Permission = "movie:delete"

	LanguageCreate Permission = "language:create"
	LanguageUpdate Permission = "language:update"
	LanguageDelete Permission = "language:delete"

	GenreCreate Permission = "genre:create"
	GenreUpdate Permission = "genre:update"
	GenreDelete Permission = "genre:delete"

	MovieFormatCreate Permission = "movie:format:create"
	MovieFormatUpdate Permission = "movie:format:update"
	MovieFormatDelete Permission = "movie:format:delete"

	PersonCreate Permission = "person:create"
	PersonUpdate Permission = "person:update"
	PersonDelete Permission = "person:delete"
)

var adminPermission = rolePermissionMap{
	MovieCreate:       true,
	MovieUpdate:       true,
	MovieDelete:       true,
	LanguageCreate:    true,
	LanguageUpdate:    true,
	LanguageDelete:    true,
	GenreCreate:       true,
	GenreUpdate:       true,
	GenreDelete:       true,
	PersonCreate:      true,
	PersonUpdate:      true,
	PersonDelete:      true,
	MovieFormatCreate: true,
	MovieFormatUpdate: true,
	MovieFormatDelete: true,
}

var regularUserPermission = rolePermissionMap{}

func (s *Auth) CheckPermissions(auth *AuthMetadata, ps ...Permission) error {
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
			return services_errors.UnauthorizedError(fmt.Errorf("user doesn't have permission %s", p))
		}
	}

	return nil
}
