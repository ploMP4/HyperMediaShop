package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"gorm.io/gorm"

	"github.com/ploMP4/HyperMediaShop/services"
	"github.com/ploMP4/HyperMediaShop/templates"
)

type ProductHandler struct {
	Logger  *httplog.Logger
	Service services.ProductService
}

// Render view that displays all products paginated
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

// Render view that displays info for a single product
func (h ProductHandler) Retrieve(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")

	product, err := h.Service.Retrieve(productID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		h.Logger.Error("error", "err", err)
		NotFoundHanlder(w, r)
		return
	} else if err != nil {
		h.Logger.Error("error", "err", err)
	}

	productPage := templates.ProductPage(product)
	err = productPage.Render(context.Background(), w)
	if err != nil {
		h.Logger.Error("error", "err", err)
	}
}
