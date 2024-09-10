package api

import (
	"context"

	"bookmymovie.app/bookmymovie"
	moviesv1 "bookmymovie.app/bookmymovie/api/gen/movies/v1"
	"bookmymovie.app/bookmymovie/services/movies/formats"
	"bookmymovie.app/bookmymovie/services/movies/genres"
	"bookmymovie.app/bookmymovie/services/movies/languages"
	"connectrpc.com/connect"
)

type moviesLanguagesService struct {
	app *bookmymovie.Application
}

func (s *moviesLanguagesService) CreateLanguage(ctx context.Context, r *connect.Request[moviesv1.CreateLanguageRequest]) (*connect.Response[moviesv1.CreateLanguageResponse], error) {
	authMeta := s.app.AuthService().GetMetadata(getAccessToken(r))
	lang, err := s.app.MoviesService().LanguagesService().Create(ctx, &authMeta, &languages.CreateParams{
		DisplayName: r.Msg.GetDisplayName(),
		EnglishName: r.Msg.GetEnglishName(),
	})
	if err != nil {
		return nil, serviceErrorHandler(err)
	}

	res := connect.NewResponse(&moviesv1.CreateLanguageResponse{
		Language: mapLanguage(&lang),
	})

	return res, nil
}

func (s *moviesLanguagesService) UpdateLanguage(ctx context.Context, r *connect.Request[moviesv1.UpdateLanguageRequest]) (*connect.Response[moviesv1.UpdateLanguageResponse], error) {
	authMeta := s.app.AuthService().GetMetadata(getAccessToken(r))
	lang, err := s.app.MoviesService().LanguagesService().Update(ctx, &authMeta, &languages.UpdateParams{
		DisplayName: r.Msg.DisplayName,
		EnglishName: r.Msg.EnglishName,
	})
	if err != nil {
		return nil, serviceErrorHandler(err)
	}

	res := connect.NewResponse(&moviesv1.UpdateLanguageResponse{
		Language: mapLanguage(&lang),
	})

	return res, nil
}

func (s *moviesLanguagesService) DeleteLanguage(ctx context.Context, r *connect.Request[moviesv1.DeleteLanguageRequest]) (*connect.Response[moviesv1.DeleteLanguageResponse], error) {
	authMeta := s.app.AuthService().GetMetadata(getAccessToken(r))
	if err := s.app.MoviesService().LanguagesService().Delete(ctx, &authMeta, r.Msg.GetId()); err != nil {
		return nil, serviceErrorHandler(err)
	}
	res := connect.NewResponse(&moviesv1.DeleteLanguageResponse{})
	return res, nil
}

func (s *moviesLanguagesService) GetLanguage(ctx context.Context, r *connect.Request[moviesv1.GetLanguageRequest]) (*connect.Response[moviesv1.GetLanguageResponse], error) {
	lang, err := s.app.MoviesService().LanguagesService().GetByID(ctx, r.Msg.GetId())
	if err != nil {
		return nil, serviceErrorHandler(err)
	}

	res := connect.NewResponse(&moviesv1.GetLanguageResponse{
		Language: mapLanguage(&lang),
	})
	return res, nil
}

func (s *moviesLanguagesService) GetLanguages(ctx context.Context, _ *connect.Request[moviesv1.GetLanguagesRequest]) (*connect.Response[moviesv1.GetLanguagesResponse], error) {
	langs, err := s.app.MoviesService().LanguagesService().List(ctx)
	if err != nil {
		return nil, serviceErrorHandler(err)
	}
	res := connect.NewResponse(&moviesv1.GetLanguagesResponse{
		Languages: mapSlice(mapLanguage, langs),
	})
	return res, nil
}

type moviesFormatsService struct {
	app *bookmymovie.Application
}

func (s *moviesFormatsService) CreateFormat(ctx context.Context, r *connect.Request[moviesv1.CreateFormatRequest]) (*connect.Response[moviesv1.CreateFormatResponse], error) {
	authMeta := s.app.AuthService().GetMetadata(getAccessToken(r))
	lang, err := s.app.MoviesService().FormatsService().Create(ctx, &authMeta, &formats.CreateParams{
		DisplayName: r.Msg.GetDisplayName(),
		About:       r.Msg.GetAbout(),
	})
	if err != nil {
		return nil, serviceErrorHandler(err)
	}

	res := connect.NewResponse(&moviesv1.CreateFormatResponse{
		Format: mapFormat(&lang),
	})

	return res, nil
}

func (s *moviesFormatsService) UpdateFormat(ctx context.Context, r *connect.Request[moviesv1.UpdateFormatRequest]) (*connect.Response[moviesv1.UpdateFormatResponse], error) {
	authMeta := s.app.AuthService().GetMetadata(getAccessToken(r))
	lang, err := s.app.MoviesService().FormatsService().Update(ctx, &authMeta, &formats.UpdateParams{
		DisplayName: r.Msg.DisplayName,
		About:       r.Msg.About,
	})
	if err != nil {
		return nil, serviceErrorHandler(err)
	}

	res := connect.NewResponse(&moviesv1.UpdateFormatResponse{
		Format: mapFormat(&lang),
	})

	return res, nil
}

func (s *moviesFormatsService) DeleteFormat(ctx context.Context, r *connect.Request[moviesv1.DeleteFormatRequest]) (*connect.Response[moviesv1.DeleteFormatResponse], error) {
	authMeta := s.app.AuthService().GetMetadata(getAccessToken(r))
	if err := s.app.MoviesService().FormatsService().Delete(ctx, &authMeta, r.Msg.GetId()); err != nil {
		return nil, serviceErrorHandler(err)
	}
	res := connect.NewResponse(&moviesv1.DeleteFormatResponse{})
	return res, nil
}

func (s *moviesFormatsService) GetFormat(ctx context.Context, r *connect.Request[moviesv1.GetFormatRequest]) (*connect.Response[moviesv1.GetFormatResponse], error) {
	lang, err := s.app.MoviesService().FormatsService().GetByID(ctx, r.Msg.GetId())
	if err != nil {
		return nil, serviceErrorHandler(err)
	}

	res := connect.NewResponse(&moviesv1.GetFormatResponse{
		Format: mapFormat(&lang),
	})
	return res, nil
}

func (s *moviesFormatsService) GetFormats(ctx context.Context, _ *connect.Request[moviesv1.GetFormatsRequest]) (*connect.Response[moviesv1.GetFormatsResponse], error) {
	langs, err := s.app.MoviesService().FormatsService().List(ctx)
	if err != nil {
		return nil, serviceErrorHandler(err)
	}
	res := connect.NewResponse(&moviesv1.GetFormatsResponse{
		Formats: mapSlice(mapFormat, langs),
	})
	return res, nil
}

type moviesGenresService struct {
	app *bookmymovie.Application
}

func (s *moviesGenresService) CreateGenre(ctx context.Context, r *connect.Request[moviesv1.CreateGenreRequest]) (*connect.Response[moviesv1.CreateGenreResponse], error) {
	authMeta := s.app.AuthService().GetMetadata(getAccessToken(r))
	lang, err := s.app.MoviesService().GenresService().Create(ctx, &authMeta, &genres.CreateParams{
		DisplayName: r.Msg.GetDisplayName(),
		About:       r.Msg.GetAbout(),
	})
	if err != nil {
		return nil, serviceErrorHandler(err)
	}

	res := connect.NewResponse(&moviesv1.CreateGenreResponse{
		Genre: mapGenre(&lang),
	})

	return res, nil
}

func (s *moviesGenresService) UpdateGenre(ctx context.Context, r *connect.Request[moviesv1.UpdateGenreRequest]) (*connect.Response[moviesv1.UpdateGenreResponse], error) {
	authMeta := s.app.AuthService().GetMetadata(getAccessToken(r))
	lang, err := s.app.MoviesService().GenresService().Update(ctx, &authMeta, &genres.UpdateParams{
		DisplayName: r.Msg.DisplayName,
		About:       r.Msg.About,
	})
	if err != nil {
		return nil, serviceErrorHandler(err)
	}

	res := connect.NewResponse(&moviesv1.UpdateGenreResponse{
		Genre: mapGenre(&lang),
	})

	return res, nil
}

func (s *moviesGenresService) DeleteGenre(ctx context.Context, r *connect.Request[moviesv1.DeleteGenreRequest]) (*connect.Response[moviesv1.DeleteGenreResponse], error) {
	authMeta := s.app.AuthService().GetMetadata(getAccessToken(r))
	if err := s.app.MoviesService().GenresService().Delete(ctx, &authMeta, r.Msg.GetId()); err != nil {
		return nil, serviceErrorHandler(err)
	}
	res := connect.NewResponse(&moviesv1.DeleteGenreResponse{})
	return res, nil
}

func (s *moviesGenresService) GetGenre(ctx context.Context, r *connect.Request[moviesv1.GetGenreRequest]) (*connect.Response[moviesv1.GetGenreResponse], error) {
	lang, err := s.app.MoviesService().GenresService().GetByID(ctx, r.Msg.GetId())
	if err != nil {
		return nil, serviceErrorHandler(err)
	}

	res := connect.NewResponse(&moviesv1.GetGenreResponse{
		Genre: mapGenre(&lang),
	})
	return res, nil
}

func (s *moviesGenresService) GetGenres(ctx context.Context, _ *connect.Request[moviesv1.GetGenresRequest]) (*connect.Response[moviesv1.GetGenresResponse], error) {
	langs, err := s.app.MoviesService().GenresService().List(ctx)
	if err != nil {
		return nil, serviceErrorHandler(err)
	}
	res := connect.NewResponse(&moviesv1.GetGenresResponse{
		Genres: mapSlice(mapGenre, langs),
	})
	return res, nil
}
