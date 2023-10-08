package main

import (
	"github.com/ploMP4/HyperMediaShop/db"
	"github.com/ploMP4/HyperMediaShop/router"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	r := router.New(db)
	r.Run()
}
