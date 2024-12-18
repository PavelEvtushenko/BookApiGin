package routes

import (
	"BookApiGin/pcg/controllers"
	"github.com/gin-gonic/gin"
)

// создаем маршруты и прописываем функции и методы для них

func RegisterBookStoreRoutesGin(r *gin.Engine) {
	r.GET("/book/", controllers.GetBook)
	r.POST("/book/", controllers.CreateBook)
	r.GET("/book/:bookId", controllers.GetBookById)
	r.PUT("/book/:bookId", controllers.UpdateBook)
	r.DELETE("/book/:bookId", controllers.DeleteBook)
	// Другие маршруты
}
