package repository

import (
	"github.com/ploMP4/HyperMediaShop/db"
	"github.com/ploMP4/HyperMediaShop/models"
)

type Product struct {
	DB *db.Database
}

func (p Product) All() ([]models.Product, error) {
	var products []models.Product

	result := p.DB.Instance.
		Preload("Reviews").
		Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (p Product) Retrieve(id string) (models.Product, error) {
	var product models.Product

	result := p.DB.Instance.
		Preload("Reviews").
		First(&product, "id = ?", id)

	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}
