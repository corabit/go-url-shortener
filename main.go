package main

import (
	"github.com/gin-gonic/gin"
	"go-url-shortener/api/controller"
	"go-url-shortener/core/repository"
	"go-url-shortener/core/services"
	"go-url-shortener/db"
	"os"
)

func main() {
	database := db.GetSqliteDatabase()
	urlRepository := repository.NewUrlRepository(database)
	urlService := services.NewUrlService(urlRepository)
	urlController := controller.NewUrlController(urlService)

	r := gin.Default()
	urlController.Route(r)
	r.Run(":8080")
}

func init() {
	err := os.Setenv("HOST", "http://localhost:8080")
	if err != nil {
		panic(err.Error())
	}
}
