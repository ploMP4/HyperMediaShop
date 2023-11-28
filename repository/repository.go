package repository

import "github.com/ploMP4/HyperMediaShop/models"

type Repository[M models.Model] interface {
	All() ([]M, error)
	Retrieve(id string) (M, error)
	// Create(form F) (T, error)
	// Update(model T) (T, error)
	// Delete(id string) (T, error)
}
