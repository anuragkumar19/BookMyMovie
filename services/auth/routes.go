package auth

import (
	"github.com/go-chi/chi/v5"
)

func (s *Auth) RegisterRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/otp", s.requestLoginOTPHandler)
		r.Post("/login", s.loginHandler)
		r.Post("/refresh", s.refreshAccessTokenHandler)
		r.Post("/logout", s.logoutHandler)
	})
}
