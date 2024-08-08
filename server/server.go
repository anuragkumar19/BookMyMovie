package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HTTPService interface {
	RegisterRoutes(chi.Router)
}

type Server struct {
	mux *chi.Mux
}

func New() Server {
	return Server{
		mux: chi.NewRouter(),
	}
}

func (srv *Server) RegisterService(s HTTPService) {
	s.RegisterRoutes(srv.mux)
}

func (srv *Server) TempServe() {
	http.ListenAndServe(":6969", srv.mux)
}
