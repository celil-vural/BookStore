package repositories

import (
	"github.com/celil-vural/BookStore/database"
	"github.com/celil-vural/BookStore/models"
)

type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
}

func GetAllBooks() ([]models.Book, error) {
	db := database.DB.Db
	var books []models.Book
	result := db.Preload("Genres").Preload("Author").Find(&books)
	return books, result.Error
}
