package handler

import (
	"github.com/celil-vural/BookStore/database"
	"github.com/celil-vural/BookStore/models"
	"github.com/celil-vural/BookStore/repositories"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAllBooks(c *fiber.Ctx) error {
	repo := repositories.BookRepository{DB: database.DB.Db}
	books, err := repo.GetAllBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get books",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"books": books})
}
func GetBookByID(c *fiber.Ctx) error {
	repo := repositories.BookRepository{DB: database.DB.Db}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id",
			"error":   err,
		})
	}
	book, err := repo.GetBookByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"book": book})
}
func CreateBook(c *fiber.Ctx) error {
	repo := repositories.BookRepository{DB: database.DB.Db}
	book := new(models.Book)
	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create book ",
			"error":   err,
		})
	}
	book, err = repo.CreateBook(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create book",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"book": book})
}
func UpdateBook(c *fiber.Ctx) error {
	repo := repositories.BookRepository{DB: database.DB.Db}
	book := new(models.Book)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id",
			"error":   err,
		})
	}
	err = c.BodyParser(book)
	if err != nil || uint(id) != book.ID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not update book",
			"error":   err,
		})
	}
	b, err := repo.GetBookByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
			"error":   err,
		})
	}
	book.CreatedAt = b.CreatedAt
	updatedBook, err := repo.UpdateBook(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update book",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"book": updatedBook})
}
func DeleteBookByID(c *fiber.Ctx) error {
	repo := repositories.BookRepository{DB: database.DB.Db}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id",
			"error":   err,
		})
	}
	book, err := repo.GetBookByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
			"error":   err,
		})
	}
	err = repo.DeleteBook(&book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete book",
			"error":   err,
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
