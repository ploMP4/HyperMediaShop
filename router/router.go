package router

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"

	"github.com/ploMP4/HyperMediaShop/db"
	"github.com/ploMP4/HyperMediaShop/handlers"
)

type Router struct {
	*chi.Mux

	db     *db.Database
	logger *httplog.Logger
}

func New(db *db.Database) *Router {
	logger := httplog.NewLogger("logger", httplog.Options{
		LogLevel:         slog.LevelDebug,
		MessageFieldName: "detail",
		Concise:          true,
	})

	r := &Router{
		Mux:    chi.NewRouter(),
		db:     db,
		logger: logger,
	}

	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode("pong")
	})

	r.NotFound(handlers.NotFoundHanlder)

	r.setupProductRoutes()
	r.setupAuthRoutes()
	r.setupAdminRoutes()

	return r
}

func (router *Router) Run() {
	_ = http.ListenAndServe(":3000", router)
}
