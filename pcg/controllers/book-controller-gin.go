package controllers

import (
	"BookApiGin/pcg/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//создаем функции для маршрутов
//*gin.Context: Это важнейший конструкт в Gin, который предоставляет доступ к запросу и позволяет отправлять ответы. В этом коде он используется для получения данных из запроса (c.BindJSON) и отправки ответов (c.JSON)
//c.JSON(http.StatusOK, newBook): Эта строка отправляет список книг в виде JSON-ответа с HTTP-статусом 200 (OK).

func GetBook(c *gin.Context) {
	newBook := models.GetAllBooks()
	c.JSON(http.StatusOK, newBook)
}

//Объявляется переменная newBook типа models.Book
//err := c.BindJSON(&newBook): Эта строка использует метод BindJSON из контекста Gin для десериализации JSON-данных из запроса в структуру newBook. Если возникает ошибка, она сохраняется в переменной err.

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

func GetBookById(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)
	c.Writer.Write(res)
}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	err = models.DeleteBook(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func UpdateBook(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	var updatedBook models.Book
	err = c.BindJSON(&updatedBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.UpdateBook(ID, &updatedBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}
