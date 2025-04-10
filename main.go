package main

import (
    "GoLangBookApi/config"
    "GoLangBookApi/models"
    "GoLangBookApi/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    config.ConnectDatabase()
    config.DB.AutoMigrate(&models.Book{})

    routes.SetupBookRoutes(app)

    app.Listen(":3000")
}
