package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string `gorm:"not null"`
	Description sql.NullString
	ImgURL      string `gorm:"default:https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg"`
	Price       int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
