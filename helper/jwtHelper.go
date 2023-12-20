package helper

import (
	"github.com/celil-vural/BookStore/config"
	"github.com/celil-vural/BookStore/models"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func AddJwt() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Config("JWT_SECRET"))},
	})
}
func CreateClaims(user *models.User) jwt.MapClaims {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		"iat":      jwt.NewNumericDate(time.Now()),
	}
	return claims
}
func GenerateJwt(user *models.User) (string, error) {
	claims := CreateClaims(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config("JWT_SECRET")))
}
