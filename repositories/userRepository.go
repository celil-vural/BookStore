package repositories

import (
	"github.com/celil-vural/BookStore/database"
	"github.com/celil-vural/BookStore/models"
)

type UserRepository interface {
	GetUsersByEmail(email string) (models.User, error)
	CreateUser(user *models.User) (models.User, error)
}

func GetUsersByEmail(email string) (models.User, error) {
	db := database.DB.Db
	var user models.User
	result := db.Where("email = ?", email).First(&user)
	return user, result.Error
}
func CreateUser(user *models.User) (models.User, error) {
	db := database.DB.Db
	result := db.Create(&user)
	return *user, result.Error
}
