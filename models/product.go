package models

import (
	"time"
)

type Product struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
