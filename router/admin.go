package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"

	"github.com/ploMP4/HyperMediaShop/middleware"
	"github.com/ploMP4/HyperMediaShop/services"
)

func (r Router) setupAdminRoutes() {
	r.Route("/admin", func(r chi.Router) {
		r.Use(jwtauth.Verifier(services.TokenAuth))
		r.Use(middleware.UnAuthenticatedRedirector)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			_, _ = w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["username"])))
		})
	})
}
