package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
	"log"
	"net/http"
	"time"

	"github.com/rokasmikas/goCAR/models"
)

func CreateOrderTable(db *pg.DB) error {

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&models.Order{}, opts)

	if err != nil {
		log.Printf("Error creating table : %v\n", err)
		return err
	}
	log.Printf("Order table created")
	return nil
}

func GetAllOrders(c *gin.Context) {
	var orders []models.Order
	carlogID := c.Param("carlogid")

	err := dbConnect.Model(&orders).Where("carlog_id = ?", carlogID).Limit(20).Select()

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
		"message": "Order data",
		"data":    orders,
	})
	return
}
// TODO: Fix those antipatterns 
func getAssociatedOrders(carID string) []models.Order{
	var orders []models.Order

	err := dbConnect.Model(&orders).Where("carlog_id = ?", carID).Select()

	if err != nil {
		log.Printf("Error: %v", err)
	}

	return orders;
}

func NewOrder(c *gin.Context) {

	var order models.Order

	c.BindJSON(&order)

	id := guuid.New().String()

	name := order.Name
	carlogid := c.Param("carlogid")

	err := dbConnect.Insert(&models.Order{
		ID:     id,
		Name:   name,
		CarlogID: carlogid,

		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
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
		"message": "Order created",
		"data": id,
	})
	return

}

// func GetCarlog(c *gin.Context) {
//
// 	carlogID := c.Param("carlogid")
// 	log.Printf("Param: %v", carlogID)
// 	carlog := &models.Carlog{ID: carlogID}
//
// 	err := dbConnect.Select(carlog)
//
// 	if err != nil {
// 		log.Printf("Error: %v", err)
//
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  http.StatusNotFound,
// 			"message": "Not found",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "Single Carlog data",
// 		"data":    carlog,
// 	})
// 	return
// }
//
// func DeleteCarlog(c *gin.Context) {
// 	carlogID := c.Param("carlogid")
//
// 	carlog := &models.Carlog{ID: carlogID}
//
// 	err := dbConnect.Delete(carlog)
//
// 	if err != nil {
// 		log.Printf("Error while deleting: %v\n", err)
//
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  http.StatusInternalServerError,
// 			"message": "Error",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "Deleted carlog successfully",
// 	})
// 	return
// }
//
// func EditCarlog(c *gin.Context) {
// 	carlogID := c.Param("carlogid")
//
// 	carlog := &models.Carlog{ID: carlogID}
// 	c.BindJSON(&carlog)
//
// 	err := dbConnect.Update(carlog)
//
// 	if err != nil {
// 		log.Printf("Error while updating: %v\n", err)
//
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  500,
// 			"message": "Error",
// 		})
//
// 	return
// 	}
//
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "Updated carlog successfully",
// 	})
// 	return
// }
