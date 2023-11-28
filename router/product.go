package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ploMP4/HyperMediaShop/handlers"
	"github.com/ploMP4/HyperMediaShop/repository"
)

func (r Router) setupProductRoutes() {
	productHandler := handlers.ProductHandler{
		Logger:     r.logger,
		Repository: repository.Product{DB: r.db},
	}

	r.Get("/", productHandler.HomePage)
	r.Route("/product", func(r chi.Router) {
		r.Get("/{id}", productHandler.ProductPage)
	})
}
