package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koralbit/go-url-shortener/api/models"
	"github.com/koralbit/go-url-shortener/core/entities"
	"github.com/koralbit/go-url-shortener/core/services"
	"net/http"
	"os"
)

type UrlController interface {
	Route(r *gin.Engine)
	Create(c *gin.Context)
	GetUrl(c *gin.Context)
}

type urlController struct {
	service services.UrlService
}

func NewUrlController(service services.UrlService) UrlController {
	return &urlController{
		service: service,
	}
}

func (c urlController) Route(r *gin.Engine) {
	r.POST("/", c.Create)
	r.GET("/:id", c.GetUrl)
}

func (c urlController) Create(ctx *gin.Context) {
	var request models.UrlCreateRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		panic(err.Error())
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
	ctx.JSON(http.StatusCreated, response)
}

func (c urlController) GetUrl(ctx *gin.Context) {
	id := ctx.Param("id")
	url := c.service.GetUrl(id)
	host := os.Getenv("HOST")
	if url == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Url not found for %s/%s", host, id),
		})
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, url.OriginUrl)
}
