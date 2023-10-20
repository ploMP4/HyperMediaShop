package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ploMP4/HyperMediaShop/db"
)

type Router struct {
	*chi.Mux

	db *db.Database
}

func New(db *db.Database) *Router {
	r := &Router{
		Mux: chi.NewRouter(),
		db:  db,
	}

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

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
