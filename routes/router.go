package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raviwu/gobookstore/controllers"
	"github.com/raviwu/gobookstore/models"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	return r
}
