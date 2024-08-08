package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Auth) RegisterRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/otp", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/refresh", func(w http.ResponseWriter, r *http.Request) {})
		r.Post("/logout", func(w http.ResponseWriter, r *http.Request) {})
	})
}
