package router

import "github.com/ploMP4/HyperMediaShop/handlers"

func (r *Router) setupProductRoutes() {
	productHandler := handlers.ProductHandler{
		DB: r.db,
	}

	r.Get("/", productHandler.All)
}
