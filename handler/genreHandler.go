package handler

import (
	"github.com/celil-vural/BookStore/models"
	"github.com/celil-vural/BookStore/repositories"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetGenres(c *fiber.Ctx) error {
	genres, err := repositories.GetAllGenres()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get genres",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"genres": genres})
}
func GetGenreByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not get genre",
			"error":   err,
		})
	}
	genre, err := repositories.GetGenreByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get genre",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"genre": genre})
}
func CreateGenre(c *fiber.Ctx) error {
	var genre models.Genre
	err := c.BodyParser(&genre)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create genre",
			"error":   err,
		})
	}
	_, err = repositories.CreateGenre(&genre)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create genre",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"genre": genre})
}
func UpdateGenre(c *fiber.Ctx) error {
	genre := new(models.Genre)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id is not valid",
			"error":   err,
		})
	}
	if err := c.BodyParser(genre); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update genre",
			"error":   err,
		})
	}
	if genre.ID == 0 || genre.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update genre",
			"error":   "ID and Name fields cannot be empty",
		})
	}
	gen, err := repositories.GetGenreByID(uint(id))
	if gen == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update genre",
			"error":   "Genre not found",
		})
	}
	genre.CreatedAt = gen.CreatedAt
	_, err = repositories.UpdateGenre(genre)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update genre",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Genre updated successfully", "genre": genre})
}
func DeleteGenreByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not delete genre",
			"error":   err,
		})
	}
	genre, err := repositories.GetGenreByID(uint(id))
	if genre == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not delete genre",
			"error":   "Genre not found",
		})
	}
	err = repositories.DeleteGenreByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete genre",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Genre deleted successfully", "genre": genre})
}
