package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ploMP4/HyperMediaShop/db"
	"github.com/ploMP4/HyperMediaShop/render"
)

type Router struct {
	*chi.Mux

	db       *db.Database
	renderer *render.Renderer
}

func New(db *db.Database) *Router {
	r := &Router{
		Mux:      chi.NewRouter(),
		db:       db,
		renderer: render.New(),
	}

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("pong")
	})

	r.setupProductRoutes()

	return r
}

func (router *Router) Run() {
	http.ListenAndServe(":3000", router)
}
