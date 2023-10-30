package router

import (
	"github.com/ploMP4/HyperMediaShop/handlers"
	"github.com/ploMP4/HyperMediaShop/services"
)

func (r Router) setupProductRoutes() {
	productHandler := handlers.ProductHandler{
		Logger:  r.logger,
		Service: services.Product{DB: r.db},
	}

	r.Get("/", productHandler.All)
}
