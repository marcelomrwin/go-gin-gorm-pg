package main

import (
	"github.com/gin-gonic/gin"
	controllers "go-gin-gorm-pg/controllers" // new
	"go-gin-gorm-pg/models"                  // new
)
func main() {
	r := gin.Default()
	db := models.SetupModels() // new
	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/books", controllers.FindBooks) // get all
	r.POST("/books", controllers.CreateBook) // create
	r.GET("/books/:id", controllers.FindBook) // find by id
	r.PATCH("/books/:id", controllers.UpdateBook) // update by id
	r.DELETE("/books/:id", controllers.DeleteBook) // delete by id
	r.Run()
}
