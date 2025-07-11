package main

import (
	//"net/http"

    "github.com/gustavobahia11/APIREST/controllers"
    "github.com/gustavobahia11/APIREST/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDB()

	router := gin.Default()

	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
	router.PATCH("/books/:id", controllers.UpdateBook)

	router.Run("localhost:8080")
}