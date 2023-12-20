package repositories

import (
	"github.com/celil-vural/BookStore/models"
	"gorm.io/gorm"
)

type GenreRepository struct {
	DB *gorm.DB
}

func (repo *GenreRepository) GetAllGenres() ([]models.Genre, error) {
	var genres []models.Genre
	result := repo.DB.Find(&genres)
	return genres, result.Error
}
func (repo *GenreRepository) GetGenreByID(id uint) (*models.Genre, error) {
	var genre models.Genre
	result := repo.DB.First(&genre, id)
	return &genre, result.Error
}
func (repo *GenreRepository) CreateGenre(genre *models.Genre) (*models.Genre, error) {
	result := repo.DB.Create(&genre)
	return genre, result.Error
}
func (repo *GenreRepository) UpdateGenre(genre *models.Genre) (*models.Genre, error) {
	result := repo.DB.Save(&genre)
	return genre, result.Error
}
func (repo *GenreRepository) DeleteGenreByID(id uint) error {
	result := repo.DB.Delete(&models.Genre{}, id)
	return result.Error
}
