package movies

import (
	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/services/movies/formats"
	"bookmymovie.app/bookmymovie/services/movies/genres"
	"bookmymovie.app/bookmymovie/services/movies/languages"
	"github.com/rs/zerolog"
)

type Movies struct {
	logger *zerolog.Logger
	db     *database.Database
	auth   *auth.Auth

	languages *languages.Languages
	formats   *formats.Formats
	genres    *genres.Genres
}

func New(logger *zerolog.Logger, db *database.Database, auth *auth.Auth) Movies {
	l := languages.New(logger, db, auth)
	f := formats.New(logger, db, auth)
	g := genres.New(logger, db, auth)
	return Movies{
		logger:    logger,
		db:        db,
		auth:      auth,
		languages: &l,
		formats:   &f,
		genres:    &g,
	}
}

func (s *Movies) LanguagesService() *languages.Languages {
	return s.languages
}

func (s *Movies) FormatsService() *formats.Formats {
	return s.formats
}
func (s *Movies) GenresService() *genres.Genres {
	return s.genres
}
