package repositories

import (
	"github.com/celil-vural/BookStore/models"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

func (repo *AuthorRepository) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	result := repo.DB.Find(&authors)
	return authors, result.Error
}
func (repo *AuthorRepository) GetAuthorByID(id int) (*models.Author, error) {
	var author models.Author
	result := repo.DB.First(&author, id)
	return &author, result.Error
}
func (repo *AuthorRepository) CreateAuthor(author *models.Author) (*models.Author, error) {
	result := repo.DB.Create(&author)
	return author, result.Error
}
func (repo *AuthorRepository) UpdateAuthor(author *models.Author) (*models.Author, error) {
	result := repo.DB.Save(&author)
	return author, result.Error
}
func (repo *AuthorRepository) DeleteAuthorByID(id uint) error {
	result := repo.DB.Delete(&models.Author{}, id)
	return result.Error
}
