package config

import (
	"github.com/go-pg/pg/v9"
	"log"
	"os"

	controllers "github.com/rokasmikas/goCAR/controllers"

)

// postgresql connection

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "test",
		Password: "test",
		Addr:     "localhost:5432",
		Database: "cargo",
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("No DB connection")
		os.Exit(100)
	}
	log.Printf("DB connection established")

	controllers.CreateCarLogTable(db)
	controllers.CreateOrderTable(db)
	controllers.InitDB(db)

	return db
}
