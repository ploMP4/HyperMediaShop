package services

import (
	"github.com/ploMP4/HyperMediaShop/db"
	"github.com/ploMP4/HyperMediaShop/models"
)

type ProductService interface {
	All() ([]models.Product, error)
	Retrieve(id string) (models.Product, error)
}

type Product struct {
	DB *db.Database
}

func (s Product) All() ([]models.Product, error) {
	var products []models.Product

	result := s.DB.Instance.
		Preload("Reviews").
		Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (s Product) Retrieve(id string) (models.Product, error) {
	var product models.Product

	result := s.DB.Instance.
		Preload("Reviews").
		First(&product, "id = ?", id)
	if result.Error != nil {
		return product, result.Error
	}

	return product, nil
}
