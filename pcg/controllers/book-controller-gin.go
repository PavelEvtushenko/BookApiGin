package controllers

import (
	"BookApiGin/pcg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//создаем функции для маршрутов

func GetBook(c *gin.Context) {
	newBook := models.GetAllBooks()
	c.JSON(http.StatusOK, newBook)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book
	err := c.BindJSON(&newBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := newBook.CreateBook()
	c.JSON(http.StatusOK, b)
}
