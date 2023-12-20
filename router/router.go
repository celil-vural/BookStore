package router

import (
	"github.com/celil-vural/BookStore/handler"
	"github.com/celil-vural/BookStore/helper"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")
	book := v1.Group("/book")
	author := v1.Group("/author")
	genre := v1.Group("/Genre")
	// routes for user
	user.Get("/login", handler.Login)
	user.Post("/register", handler.Register)
	// routes for book
	book.Get("/", handler.GetAllBooks)
	book.Get("/:id", handler.GetBookByID)
	book.Post("/", handler.CreateBook)
	book.Put("/:id", handler.UpdateBook)
	book.Delete("/:id", handler.DeleteBookByID)
	// routes for author
	author.Get("/", handler.GetAuthors)
	author.Get("/:id", handler.GetAuthorByID)
	author.Post("/", helper.AddJwt(), handler.CreateAuthor)
	author.Put("/:id", helper.AddJwt(), handler.UpdateAuthor)
	author.Delete("/:id", helper.AddJwt(), handler.DeleteAuthorByID)
	// routes for genre
	genre.Get("/", handler.GetGenres)
	genre.Get("/:id", handler.GetGenreByID)
	genre.Post("/", helper.AddJwt(), handler.CreateGenre)
	genre.Put("/:id", helper.AddJwt(), handler.UpdateGenre)
	genre.Delete("/:id", helper.AddJwt(), handler.DeleteGenreByID)
}
