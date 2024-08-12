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
)

var adminPermission = rolePermissionMap{
	MovieCreate: true,
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
