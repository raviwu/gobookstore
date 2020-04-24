package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/raviwu/gobookstore/models"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
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

func FindBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		fmt.Printf("request book id [%s] does not exist", c.Param("id"))
		c.AbortWithStatus(http.StatusOK)
		return
	}

	db.Delete(&book)
	c.AbortWithStatus(http.StatusOK)
}
