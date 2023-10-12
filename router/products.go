package router

import "github.com/ploMP4/HyperMediaShop/handlers"

func (r *Router) setupProductRoutes() {
	productHandler := handlers.ProductHandler{
		DB:       r.db,
		Renderer: r.renderer,
	}

	r.Get("/", productHandler.All)
}
