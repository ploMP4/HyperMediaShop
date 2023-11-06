package models

import (
	"database/sql"
	"time"
)

type Review struct {
	UserID    string `gorm:"primaryKey;type:uuid"`
	ProductID string `gorm:"primaryKey;type:uuid"`

	Title       string `gorm:"not null;type:varchar(150)"`
	Description sql.NullString
	Rating      int `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
