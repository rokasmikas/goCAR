package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type Carlog struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
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

func GetAllCarlogs(c *gin.Context) {
	var carlogs []Carlog

	err := dbConnect.Model(&carlogs).Select()

	if err != nil {
		log.Printf("Error: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Carlog data",
		"data":    carlogs,
	})
	return
}

func NewCarlog(c *gin.Context) {
	log.Printf("Init")
	var carlog Carlog
	c.BindJSON(&carlog)

	id := guuid.New().String()

	name := carlog.Name
	make := carlog.Make
	model := carlog.Model
	reg := carlog.Reg
	year := carlog.Year
	active := "True"

	// TODO: Sanitize inputs
	err := dbConnect.Insert(&Carlog{
		ID:     id,
		Name:   name,
		Make:   make,
		Model:  model,
		Reg:    reg,
		Year:   year,
		Active: active,
	})

	if err != nil {
		log.Printf("Error while creating: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Carlog created",
	})
	return

}
