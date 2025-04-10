package routes

import (
	"GoLangBookApi/handlers"
	"github.com/gofiber/fiber/"
)

func SetupBookRoutes(app *fiber.App) {
    api := app.Group("/books")
    api.Post("/", handlers.CreateBook)
    api.Get("/", handlers.GetBooks)
    api.Get("/:id", handlers.GetBook)
    api.Put("/:id", handlers.UpdateBook)
    api.Delete("/:id", handlers.DeleteBook)
}