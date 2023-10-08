package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ploMP4/HyperMediaShop/models"
)

type Database struct {
	Conn *gorm.DB
}

func Connect() (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db := &Database{
		Conn: conn,
	}

	return db, nil
}

func (db *Database) Migrate() error {
	return db.Conn.AutoMigrate(&models.Product{})
}
