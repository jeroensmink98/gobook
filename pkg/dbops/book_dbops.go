package dbops

import (
	"github.com/jeroensmink98/gobook/pkg/models"
	"gorm.io/gorm"
)

func GetAllBooks(db *gorm.DB) ([]models.Book, error) {
	var books []models.Book
	res := db.Find(&books)
	if res.Error != nil {
		return nil, res.Error
	}
	return books, nil
}

func GetBookByID(db *gorm.DB, id uint) (*models.Book, error) {
	var book models.Book
	res := db.First(&book, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &book, nil
}

func CreateNewBook(db *gorm.DB, book *models.Book) error {
	res := db.Create(book)
	return res.Error
}

func DeleteBook(db *gorm.DB, id uint) (*models.Book, error) {
	var book models.Book
	res := db.Delete(&book, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &book, nil
}
