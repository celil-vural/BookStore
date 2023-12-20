package repositories

import (
	"github.com/celil-vural/BookStore/database"
	"github.com/celil-vural/BookStore/models"
)

type GenreRepository interface {
	GetAllGenres() ([]models.Genre, error)
	CreateGenre(genre *models.Genre) (*models.Genre, error)
	UpdateGenre(genre *models.Genre) (*models.Genre, error)
	GetGenreByID(id int) (*models.Genre, error)
	DeleteGenreByID(id int) error
}

func GetAllGenres() ([]models.Genre, error) {
	db := database.DB.Db
	var genres []models.Genre
	result := db.Find(&genres)
	return genres, result.Error
}
func GetGenreByID(id uint) (*models.Genre, error) {
	db := database.DB.Db
	var genre models.Genre
	result := db.First(&genre, id)
	return &genre, result.Error
}
func CreateGenre(genre *models.Genre) (*models.Genre, error) {
	db := database.DB.Db
	result := db.Create(&genre)
	return genre, result.Error
}
func UpdateGenre(genre *models.Genre) (*models.Genre, error) {
	db := database.DB.Db
	result := db.Save(&genre)
	return genre, result.Error
}
func DeleteGenreByID(id uint) error {
	db := database.DB.Db
	result := db.Delete(&models.Genre{}, id)
	return result.Error
}
