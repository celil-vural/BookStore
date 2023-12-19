package router

import "github.com/gofiber/fiber/v2"

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/v1")
	book := v1.Group("/book")
	author := v1.Group("/author")
	Genre := v1.Group("/Genre")
	// routes for book
	book.Get("book", handler.GetAllBooks)
	book.Get("book/:id", handler.GetBookByID)
	book.Post("book", handler.CreateBook)
	book.Put("book/:id", handler.UpdateBook)
	book.Delete("book/:id", handler.DeleteBookByID)
	book.Get("/search/:query", handler.SearchBooks)
	book.Get("author/:author", handler.GetBooksByAuthor)
	book.Get("genre/:genre", handler.GetBooksByGenre)
	// routes for author
	author.Get("/", handler.GetAuthors)
	author.Get("/:id", handler.GetAuthorByID)
	author.Post("/", handler.CreateAuthor)
	author.Put("/:id", handler.UpdateAuthor)
	author.Delete("/:id", handler.DeleteAuthorByID)
	// routes for genre
	Genre.Get("/", handler.GetGenres)
	Genre.Get("/:id", handler.GetGenreByID)
	Genre.Post("/", handler.CreateGenre)
	Genre.Put("/:id", handler.UpdateGenre)
	Genre.Delete("/:id", handler.DeleteGenreByID)

}
