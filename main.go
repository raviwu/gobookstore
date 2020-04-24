package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raviwu/gobookstore/controllers"
	"github.com/raviwu/gobookstore/models"
)

func main() {
	r := gin.Default()

	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)

	r.Run()
}
