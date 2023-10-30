package services

import (
	"github.com/ploMP4/HyperMediaShop/db"
	"github.com/ploMP4/HyperMediaShop/models"
)

type ProductService interface {
	All() ([]models.Product, error)
}

type Product struct {
	DB *db.Database
}

func (s Product) All() ([]models.Product, error) {
	var products []models.Product

	result := s.DB.Instance.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
