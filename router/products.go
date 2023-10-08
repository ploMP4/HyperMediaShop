package router

import "github.com/ploMP4/HyperMediaShop/handlers"

func (r *Router) setupProductRoutes() {
	handler := handlers.ProductHandler{DB: r.db}

	r.Get("/product", handler.All)
}
