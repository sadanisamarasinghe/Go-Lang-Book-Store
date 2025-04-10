package handlers

import (
	"GoLangBookApi/config"
	"GoLangBookApi/models"
	"github.com/gofiber/fiber/v2"
)

//create a new book
func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&Book)
	return c.Status(201).JSON(book)
}

//Get all books
func GetBooks(c *fiber.Ctx) error{
	var books []models.Book
	config.DB.Find(&books)
	return c.JSON(books)
}

//Get book by ID
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := config.DB.First(&book, id).Error; err !=nil{
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

//update a book
func UpdateBook(c *fiber.Ctx) error{
	id := c.Params("id")
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil{
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})

	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
	config.DB.Save(&book)
	return c.JSON(book)
	}
}

//delete a book
func DeleteBook(c *fiber.Ctx) error{
	id := c.Params("id")
	result := config.DB.Delete(&models.Book{}, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(204)
	}
}