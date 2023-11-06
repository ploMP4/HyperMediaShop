package models

type User struct {
	BaseModel

	Username  string `gorm:"not null;type:varchar(150);uniqueIndex"`
	Email     string `gorm:"not null;type:varchar(150);uniqueIndex"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"type:varchar(150)"`
	LastName  string `gorm:"type:varchar(150)"`
	Phone     string `gorm:"type:varchar(150)"`
	Country   string `gorm:"type:varchar(150)"`
	City      string `gorm:"type:varchar(150)"`
	Address   string `gorm:"type:varchar(150)"`
	IsStaff   bool   `gorm:"default:false"`

	Orders  []Order
	Reviews []Review
}
