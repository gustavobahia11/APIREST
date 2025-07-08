package controllers

import (
	"Users/jader/APIREST/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(ctx *gin.Context) {
	var book models.Book

	err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(ctx *gin.Context) {
	var input CreateBookInput

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(ctx *gin.Context) {
	var book models.Book

	err := models.DB.Where("id = ?", ctx.Param("id")).Delete(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data deleted": book})
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func UpdateBook(ctx *gin.Context) {
	var book models.Book

	err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var input UpdateBookInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	models.DB.Model(&book).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"updated": book})
}