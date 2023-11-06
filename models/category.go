package models

type Category struct {
	BaseModel

	Name string `gorm:"not null;type:varchar(150)"`

	Products []*Product `gorm:"many2many:product_category"`
}
