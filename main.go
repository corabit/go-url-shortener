package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/koralbit/go-url-shortener/api/controller"
	"github.com/koralbit/go-url-shortener/core/repository"
	"github.com/koralbit/go-url-shortener/core/services"
	"github.com/koralbit/go-url-shortener/db"
	"github.com/labstack/echo/v4"
)

func main() {
	database := db.GetPostgreSQL(getDBConfig())
	urlRepository := repository.NewUrlRepository(database)
	urlService := services.NewUrlService(urlRepository)
	urlController := controller.NewUrlController(urlService)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	urlController.Route(e)
	for _, r := range e.Routes() {
		fmt.Println(r.Path)
	}
	e.Start(":8080")
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func getDBConfig() db.PostgreSQLConfiguration {
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	return db.PostgreSQLConfiguration{
		User:     db_user,
		Password: db_pass,
		Host:     db_host,
		Port:     db_port,
		Database: db_name,
	}
}

func init() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	err := os.Setenv("HOST", fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		panic(err.Error())
	}
}
