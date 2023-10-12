package handlers

import (
	"net/http"

	"github.com/ploMP4/HyperMediaShop/db"
	"github.com/ploMP4/HyperMediaShop/render"
)

type ProductHandler struct {
	DB       *db.Database
	Renderer *render.Renderer
}

func (h ProductHandler) All(w http.ResponseWriter, _ *http.Request) {
	h.Renderer.ExecuteTemplate(w, "home", nil)
}
