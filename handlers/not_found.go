package handlers

import (
	"context"
	"net/http"

	"github.com/ploMP4/HyperMediaShop/templates"
)

func NotFoundHanlder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	NotFound := templates.NotFound()
	_ = NotFound.Render(context.Background(), w)
}
