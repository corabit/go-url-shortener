package services

import (
	"errors"
	"github.com/koralbit/go-url-shortener/core/entities"
	"github.com/koralbit/go-url-shortener/core/repository"
	"github.com/lithammer/shortuuid/v3"
	"gorm.io/gorm"
)

type UrlService interface {
	Create(url entities.Url) entities.Url
	GetUrl(id string) *entities.Url
}

type urlService struct {
	repo repository.UrlRepository
}

func NewUrlService(urlRepository repository.UrlRepository) UrlService {
	return &urlService{
		repo: urlRepository,
	}
}

func (s urlService) Create(url entities.Url) entities.Url {
	url.Id = shortuuid.New()
	u, err := s.repo.Insert(url)
	if err != nil {
		panic(err)
	}
	return u
}

func (s urlService) GetUrl(id string) *entities.Url {
	u, err := s.repo.FindById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		panic(err.Error())
	}
	u.Clicks++
	u, err = s.repo.Update(u)
	if err != nil {
		panic(err.Error())
	}
	return &u
}
