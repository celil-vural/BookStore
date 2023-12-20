package repositories

import (
	"github.com/celil-vural/BookStore/database"
	"github.com/celil-vural/BookStore/models"
)

type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (models.Book, error)
	CreateBook(book *models.Book) (*models.Book, error)
	UpdateBook(book models.Book) (*models.Book, error)
	DeleteBook(book models.Book) error
}

func GetAllBooks() ([]models.Book, error) {
	db := database.DB.Db
	var books []models.Book
	result := db.Preload("Genres").Preload("Author").Find(&books)
	return books, result.Error
}
func GetBookByID(id uint) (models.Book, error) {
	db := database.DB.Db
	var book models.Book
	result := db.Preload("Genres").Preload("Author").First(&book, id)
	return book, result.Error
}
func CreateBook(book *models.Book) (*models.Book, error) {
	db := database.DB.Db
	result := db.Create(&book)
	return book, result.Error
}
func UpdateBook(book *models.Book) (*models.Book, error) {
	db := database.DB.Db
	result := db.Save(&book)
	return book, result.Error
}
func DeleteBook(book *models.Book) error {
	db := database.DB.Db
	result := db.Delete(&book)
	return result.Error
}
