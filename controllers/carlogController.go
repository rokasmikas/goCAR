package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
	"log"
	"net/http"

	"github.com/rokasmikas/goCAR/models"
)

func CreateCarLogTable(db *pg.DB) error {

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&models.Carlog{}, opts)

	if err != nil {
		log.Printf("Error creating table : %v\n", err)
		return err
	}
	log.Printf("Carlog table created")
	return nil
}

func GetAllCarlogs(c *gin.Context) {
	var carlogs []models.Carlog

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

	var carlog models.Carlog

	c.BindJSON(&carlog)

	id := guuid.New().String()

	name := carlog.Name
	make := carlog.Make
	model := carlog.Model
	reg := carlog.Reg
	year := carlog.Year
	active := "True"

	// TODO: Sanitize inputs
	err := dbConnect.Insert(&models.Carlog{
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
		"data": id,
	})
	return

}

func GetCarlog(c *gin.Context) {

	carlogID := c.Param("carlogid")
	log.Printf("Param: %v", carlogID)
	carlog := &models.Carlog{ID: carlogID}

	err := dbConnect.Select(carlog)

	if err != nil {
		log.Printf("Error: %v", err)

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Carlog data",
		"data":    carlog,
	})
	return
}

func DeleteCarlog(c *gin.Context) {
	carlogID := c.Param("carlogid")

	carlog := &models.Carlog{ID: carlogID}

	err := dbConnect.Delete(carlog)

	if err != nil {
		log.Printf("Error while deleting: %v\n", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted carlog successfully",
	})
	return
}

func EditCarlog(c *gin.Context) {
	carlogID := c.Param("carlogid")

	carlog := &models.Carlog{ID: carlogID}
	c.BindJSON(&carlog)

	err := dbConnect.Update(carlog)

	if err != nil {
		log.Printf("Error while updating: %v\n", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error",
		})

	return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Updated carlog successfully",
	})
	return
}
