package handler

import (
	"github.com/celil-vural/BookStore/database"
	"github.com/celil-vural/BookStore/models"
	"github.com/celil-vural/BookStore/repositories"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAuthors(c *fiber.Ctx) error {
	repo := repositories.AuthorRepository{DB: database.DB.Db}
	authors, err := repo.GetAllAuthors()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get authors",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"authors": authors})
}
func CreateAuthor(c *fiber.Ctx) error {
	repo := repositories.AuthorRepository{DB: database.DB.Db}
	author := new(models.Author)
	if err := c.BodyParser(author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create author",
			"error":   err,
		})
	}
	_, err := repo.CreateAuthor(author)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create author",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Author created successfully", "author": author})
}
func UpdateAuthor(c *fiber.Ctx) error {
	repo := repositories.AuthorRepository{DB: database.DB.Db}
	author := new(models.Author)
	id, err := strconv.Atoi(c.Params("id"))
	if err := c.BodyParser(author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update author",
			"error":   err,
		})
	}
	if err != nil || author.ID != uint(id) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update author",
			"error":   "ID is not valid",
		})
	}
	if author.ID == 0 || author.Name == "" || author.Surname == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update author",
			"error":   "ID, Name and Surname fields cannot be empty",
		})
	}
	auth, err := repo.GetAuthorByID(id)
	if auth == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update author",
			"error":   "Author not found",
		})
	} else if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update author",
			"error":   err,
		})
	}
	author.CreatedAt = auth.CreatedAt
	_, err = repo.UpdateAuthor(author)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update author",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Author updated successfully", "author": author})
}
func GetAuthorByID(c *fiber.Ctx) error {
	repo := repositories.AuthorRepository{DB: database.DB.Db}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not get author",
			"error":   err,
		})
	}
	author, err := repo.GetAuthorByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get author",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"author": author})
}
func DeleteAuthorByID(c *fiber.Ctx) error {
	repo := repositories.AuthorRepository{DB: database.DB.Db}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not delete author",
			"error":   err,
		})
	}
	author, err := repo.GetAuthorByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete author",
			"error":   err,
		})
	}
	err = repo.DeleteAuthorByID(author.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete author",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Author deleted successfully", "author": author})
}
