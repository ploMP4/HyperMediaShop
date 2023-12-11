package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"

	"github.com/ploMP4/HyperMediaShop/handlers"
	"github.com/ploMP4/HyperMediaShop/middleware"
	"github.com/ploMP4/HyperMediaShop/services"
)

func (r Router) setupAuthRoutes() {
	authHandler := handlers.AuthHandler{
		Logger: r.logger,
	}

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(services.TokenAuth))
		r.Use(middleware.AuthenticatedRedirector)

		r.Post("/login", authHandler.Login)
		r.Get("/login", authHandler.LoginPage)

		r.Post("/register", authHandler.Register)
		r.Get("/register", authHandler.RegisterPage)
	})

	r.Get("/logout", authHandler.Logout)
	r.Post("/logout", authHandler.Logout)
}
