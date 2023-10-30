package router

import "github.com/ploMP4/HyperMediaShop/handlers"

func (r *Router) setupProductRoutes() {
	productHandler := handlers.ProductHandler{
		DB:     r.db,
		Logger: r.logger,
	}

	r.Get("/", productHandler.All)
}
