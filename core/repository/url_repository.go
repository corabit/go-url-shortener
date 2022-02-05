package repository

import (
	"github.com/koralbit/go-url-shortener/core/entities"
	"gorm.io/gorm"
)

type UrlRepository interface {
	Insert(url entities.Url) (entities.Url, error)
	FindById(id string) (entities.Url, error)
	Update(url entities.Url) (entities.Url, error)
}

type urlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	err := db.AutoMigrate(&entities.Url{})
	if err != nil {
		panic(err.Error())
	}
	return &urlRepository{
		db: db,
	}
}

func (u urlRepository) Insert(url entities.Url) (entities.Url, error) {
	err := u.db.Create(&url).Error
	return url, err
}

func (u urlRepository) FindById(id string) (entities.Url, error) {
	var url entities.Url
	err := u.db.First(&url, "id = ?", id).Error
	return url, err
}

func (u urlRepository) Update(url entities.Url) (entities.Url, error) {
	err := u.db.Updates(&url).Error
	return url, err
}
