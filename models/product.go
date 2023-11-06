package models

import "database/sql"

type Product struct {
	BaseModel

	Name        string `gorm:"not null;type:varchar(150)"`
	Description sql.NullString
	ImgURL      string `gorm:"default:https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg"`
	Price       int    `gorm:"not null"`

	Categories []*Category `gorm:"many2many:product_category"`
	Reviews    []Review
}
