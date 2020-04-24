package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/raviwu/gobookstore/models"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(c *gin.Context) {
	// Validate Input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	db := c.MustGet("db").(*gorm.DB)

	book := models.Book{Title: input.Title, Author: input.Author}
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var books []models.Book
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
