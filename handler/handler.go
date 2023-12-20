package handler

import (
	"github.com/celil-vural/BookStore/helper"
	"github.com/celil-vural/BookStore/models"
	"github.com/celil-vural/BookStore/repositories"
	"github.com/gofiber/fiber/v2"
)

// Login user
func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := repositories.GetUsersByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err,
		})
	}
	validate := helper.ValidatePassword(user.Password, password)
	if !validate {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Wrong password",
		})
	}
	token, err := helper.GenerateJwt(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{"token": token})
}

// Register user
func Register(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	username := c.FormValue("username")
	_, err := repositories.GetUsersByEmail(email)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}
	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not register user",
			"error":   err,
		})
	}
	newUser := models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
	_, err = repositories.CreateUser(&newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not register user",
			"error":   err,
		})
	}
	token, err := helper.GenerateJwt(&newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
func GetAllBooks(c *fiber.Ctx) error {
	books, err := repositories.GetAllBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get books",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"books": books})
}
