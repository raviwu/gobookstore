package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raviwu/gobookstore/models"
)

func main() {
	r := gin.Default()

	models.SetupModels()

	r.Run()
}
