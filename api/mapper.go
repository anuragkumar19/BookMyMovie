package api

import (
	"time"

	moviesv1 "bookmymovie.app/bookmymovie/api/gen/movies/v1"
	usersv1 "bookmymovie.app/bookmymovie/api/gen/users/v1"
	"bookmymovie.app/bookmymovie/database"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapRole(role database.Roles) usersv1.Role {
	switch role {
	case database.RolesAdmin:
		return usersv1.Role_ROLE_ADMIN
	default:
		return usersv1.Role_ROLE_REGULAR_USER
	}
}

func mapLanguage(lang *database.MoviesLanguage) *moviesv1.Language {
	return &moviesv1.Language{
		Id:          lang.ID,
		DisplayName: lang.DisplayName,
		EnglishName: lang.EnglishName,
		Slug:        lang.Slug,
	}
}

func mapFormat(format *database.MoviesFormat) *moviesv1.Format {
	return &moviesv1.Format{
		Id:          format.ID,
		DisplayName: format.DisplayName,
		About:       format.About,
		Slug:        format.Slug,
	}
}

func mapGenre(genre *database.MoviesGenre) *moviesv1.Genre {
	return &moviesv1.Genre{
		Id:          genre.ID,
		DisplayName: genre.DisplayName,
		About:       genre.About,
		Slug:        genre.Slug,
	}
}

func mapSlice[T any, S any](f func(*T) S, s []T) []S {
	ss := make([]S, 0, len(s))
	for _, l := range s {
		ss = append(ss, f(&l))
	}
	return ss
}

func mapDate(t time.Time, valid bool) *date.Date {
	if !valid {
		return nil
	}

	return &date.Date{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}

func mapUser(user *database.FindUserByIdRow) *usersv1.User {
	return &usersv1.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Version:   user.Version,
		Role:      mapRole(user.Role),
		Dob:       mapDate(user.Dob.Time, user.Dob.Valid),
		CreatedAt: timestamppb.New(user.CreatedAt.Time),
	}
}
