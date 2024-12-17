package main

import (
	"BookApiGin/pcg/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// запускаем сервер и определяем маршруты
// r := gin.Default() создает новый экземпляр Gin роутера с включенным средним уровнем логирования и рекавери (восстановление после паники)

func main() {
	r := gin.Default()
	routes.RegisterBookStoreRoutesGin(r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
