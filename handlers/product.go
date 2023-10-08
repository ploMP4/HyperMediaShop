package handlers

import (
	"net/http"

	"github.com/ploMP4/HyperMediaShop/db"
)

type ProductHandler struct {
	DB *db.Database
}

func (h ProductHandler) All(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello Products"))
}
