package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/httplog/v2"

	"github.com/ploMP4/HyperMediaShop/services"
	"github.com/ploMP4/HyperMediaShop/templates"
)

type ProductHandler struct {
	Logger  *httplog.Logger
	Service services.ProductService
}

func (h ProductHandler) All(w http.ResponseWriter, r *http.Request) {
	products, err := h.Service.All()
	if err != nil {
		h.Logger.Error("error", "err", err)
	}

	homePage := templates.Index(products)
	err = homePage.Render(context.Background(), w)
	if err != nil {
		h.Logger.Error("error", "err", err)
	}
}
