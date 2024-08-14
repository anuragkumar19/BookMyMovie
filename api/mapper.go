package api

import (
	usersv1 "bookmymovie.app/bookmymovie/api/gen/users/v1"
	"bookmymovie.app/bookmymovie/database"
)

func mapRole(role database.Roles) usersv1.Role {
	switch role {
	case database.RolesAdmin:
		return usersv1.Role_ROLE_ADMIN
	default:
		return usersv1.Role_ROLE_REGULAR_USER
	}
}
