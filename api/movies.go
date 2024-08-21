package api

import (
	"bookmymovie.app/bookmymovie/api/gen/movies/v1/moviesv1connect"
)

type moviesLanguagesService struct {
	moviesv1connect.UnimplementedMoviesLanguagesServiceHandler
}

type moviesGenresService struct {
	moviesv1connect.UnimplementedMoviesGenresServiceHandler
}

type moviesFormatsService struct {
	moviesv1connect.UnimplementedMoviesFormatsServiceHandler
}
