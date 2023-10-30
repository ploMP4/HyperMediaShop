package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/httplog/v2"

	"github.com/ploMP4/HyperMediaShop/db"
	"github.com/ploMP4/HyperMediaShop/models"
	"github.com/ploMP4/HyperMediaShop/templates"
)

type ProductHandler struct {
	DB     *db.Database
	Logger *httplog.Logger
}

func (h ProductHandler) All(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	result := h.DB.Instance.Find(&products)
	if result.Error != nil {
		h.Logger.Error("error", "err", result.Error)
	}

	homePage := templates.Index(products)
	err := homePage.Render(context.Background(), w)
	if err != nil {
		h.Logger.Error("error", "err", err)
	}
}
