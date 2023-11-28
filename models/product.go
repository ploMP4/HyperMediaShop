package models

import (
	"database/sql"
)

type Product struct {
	BaseModel

	Name           string `gorm:"not null;type:varchar(150)"`
	Description    sql.NullString
	PreviewImgURL  string `gorm:"default:https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg"`
	DisplayImgURL1 string `gorm:"default:https://tailwindui.com/img/ecommerce-images/product-page-02-secondary-product-shot.jpg"`
	DisplayImgURL2 string `gorm:"default:https://tailwindui.com/img/ecommerce-images/product-page-02-tertiary-product-shot-01.jpg"`
	DisplayImgURL3 string `gorm:"default:https://tailwindui.com/img/ecommerce-images/product-page-02-tertiary-product-shot-02.jpg"`
	DisplayImgURL4 string `gorm:"default:https://tailwindui.com/img/ecommerce-images/product-page-02-featured-product-shot.jpg"`
	Price          int    `gorm:"not null"`

	Categories []*Category `gorm:"many2many:product_category"`
	Reviews    []Review
}

func (p Product) String() string {
	return p.Name
}
