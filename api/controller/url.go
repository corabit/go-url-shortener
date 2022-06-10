package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/koralbit/go-url-shortener/api/models"
	"github.com/koralbit/go-url-shortener/core/entities"
	"github.com/koralbit/go-url-shortener/core/services"
	"github.com/labstack/echo/v4"
)

type UrlController interface {
	Route(e *echo.Echo)
	Create(c echo.Context) error
	GetUrl(c echo.Context) error
}

type urlController struct {
	service services.UrlService
}

func NewUrlController(service services.UrlService) UrlController {

	return &urlController{
		service: service,
	}
}

func (c urlController) Route(e *echo.Echo) {
	e.POST("", c.Create)
	e.GET("/:id", c.GetUrl)
}

func (c urlController) Create(ctx echo.Context) error {
	var request models.UrlCreateRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(request); err != nil {
		return err
	}
	url := entities.Url{
		OriginUrl: request.Url,
	}
	curl := c.service.Create(url)
	host := os.Getenv("HOST")
	response := models.UrlCreateResponse{
		Id:       curl.Id,
		ShortUrl: fmt.Sprintf("%s/%s", host, curl.Id),
	}
	return ctx.JSON(http.StatusCreated, response)
}

func (c urlController) GetUrl(ctx echo.Context) error {
	id := ctx.Param("id")
	url := c.service.GetUrl(id)
	host := os.Getenv("HOST")
	if url == nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"error": fmt.Sprintf("Url not found for %s/%s", host, id),
		})
	}
	return ctx.Redirect(http.StatusMovedPermanently, url.OriginUrl)
}
