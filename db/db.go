package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ploMP4/HyperMediaShop/models"
)

type Database struct {
	Instance *gorm.DB
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
		Instance: conn,
	}

	return db, nil
}

func (db *Database) Migrate() error {
	err := db.Instance.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	err = db.Instance.AutoMigrate(&models.Product{})
	if err != nil {
		return err
	}

	err = db.Instance.AutoMigrate(&models.Category{})
	if err != nil {
		return err
	}

	err = db.Instance.AutoMigrate(&models.Review{})
	if err != nil {
		return err
	}

	err = db.Instance.AutoMigrate(&models.Order{})
	if err != nil {
		return err
	}

	return nil
}
