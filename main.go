package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koralbit/go-url-shortener/api/controller"
	"github.com/koralbit/go-url-shortener/core/repository"
	"github.com/koralbit/go-url-shortener/core/services"
	"github.com/koralbit/go-url-shortener/db"
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
