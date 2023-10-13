package handlers

import (
	"context"
	"net/http"

	"github.com/ploMP4/HyperMediaShop/db"
	"github.com/ploMP4/HyperMediaShop/templates"
)

type ProductHandler struct {
	DB *db.Database
}

func (h ProductHandler) All(w http.ResponseWriter, _ *http.Request) {
	homePage := templates.Index()
	homePage.Render(context.Background(), w)
}
