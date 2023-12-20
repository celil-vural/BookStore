package repositories

import (
	"github.com/celil-vural/BookStore/database"
	"github.com/celil-vural/BookStore/models"
)

type AuthorRepository interface {
	GetAllAuthors() ([]models.Author, error)
	CreateAuthor(author models.Author) (*models.Author, error)
	UpdateAuthor(author models.Author) (*models.Author, error)
	GetAuthorByID(id int) (*models.Author, error)
	DeleteAuthorByID(id int) error
}

func GetAllAuthors() ([]models.Author, error) {
	db := database.DB.Db
	var authors []models.Author
	result := db.Find(&authors)
	return authors, result.Error
}
func GetAuthorByID(id int) (*models.Author, error) {
	db := database.DB.Db
	var author models.Author
	result := db.First(&author, id)
	return &author, result.Error
}
func CreateAuthor(author *models.Author) (*models.Author, error) {
	db := database.DB.Db
	result := db.Create(&author)
	return author, result.Error
}
func UpdateAuthor(author *models.Author) (*models.Author, error) {
	db := database.DB.Db
	result := db.Save(&author)
	return author, result.Error
}
func DeleteAuthorByID(id uint) error {
	db := database.DB.Db
	result := db.Delete(&models.Author{}, id)
	return result.Error
}
