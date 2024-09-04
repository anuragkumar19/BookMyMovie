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

func New(logger *zerolog.Logger, db *database.Database, a *auth.Auth) Movies {
	l := languages.New(logger, db, a)
	f := formats.New(logger, db, a)
	g := genres.New(logger, db, a)
	return Movies{
		logger:    logger,
		db:        db,
		auth:      a,
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
