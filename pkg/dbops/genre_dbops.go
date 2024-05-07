package dbops

import (
	"github.com/jeroensmink98/gobook/pkg/models"
	"gorm.io/gorm"
)

func GetAllGenres(db *gorm.DB) ([]models.Genre, error) {
	var genres []models.Genre
	res := db.Find(&genres)
	if res.Error != nil {
		return nil, res.Error
	}
	return genres, nil
}

func GetGenreByID(db *gorm.DB, id uint) (*models.Genre, error) {
	var genre models.Genre
	res := db.First(&genre, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &genre, nil
}

func CreateGenre(db *gorm.DB, genre *models.Genre) error {
	res := db.Create(genre)
	return res.Error
}
