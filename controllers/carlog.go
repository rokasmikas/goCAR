package controllers

import (
	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"

	"log"
	"time"
)

type Carlog struct {
	ID        string    `json:"id"`
	Make      string    `json:"make"`
	Model     string    `json:"model"`
	Reg       string    `json:"reg"`
	Year      string    `json:"year"`
	Active    string    `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



var dbConnect *pg.DB

func InitDB(db *pg.DB) {
  dbConnect = db
}



func CreateCarLogTable(db *pg.DB) error {

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&Carlog{}, opts)

	if err != nil {
		log.Printf("Error creating table : %v\n", err)
		return err
	}
	log.Printf("Carlog table created")
	return nil
}
