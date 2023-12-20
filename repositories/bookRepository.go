package repositories

import (
	"github.com/celil-vural/BookStore/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (repo *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := repo.DB.Preload("Genres").Preload("Author").Find(&books)
	return books, result.Error
}
func (repo *BookRepository) GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	result := repo.DB.Preload("Genres").Preload("Author").First(&book, id)
	return book, result.Error
}
func (repo *BookRepository) CreateBook(book *models.Book) (*models.Book, error) {
	result := repo.DB.Create(&book)
	return book, result.Error
}
func (repo *BookRepository) UpdateBook(book *models.Book) (*models.Book, error) {
	result := repo.DB.Preload("Genres").Preload("Author").Save(&book)
	return book, result.Error
}
func (repo *BookRepository) DeleteBook(book *models.Book) error {
	result := repo.DB.Delete(&book)
	return result.Error
}
