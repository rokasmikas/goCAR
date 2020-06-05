package controllers

import (
	"github.com/go-pg/pg/v9"
)

var dbConnect *pg.DB

func InitDB(db *pg.DB) {
	dbConnect = db
}
