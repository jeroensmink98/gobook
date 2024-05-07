package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jeroensmink98/gobook/pkg/dbops"
	"github.com/jeroensmink98/gobook/pkg/models"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		books, err := dbops.GetAllBooks(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error retrieving books: %v", err))
		}
		return c.JSON(books)
	}
}

func GetBookById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("invalid id")
		}

		book, err := dbops.GetBookByID(db, uint(id))
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("book with id %v not found", id))
		}
		return c.JSON(book)
	}
}

func CreateBook(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var book models.Book

		if err := c.BodyParser(&book); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("parse error: %v", err))
		}

		err := dbops.CreateNewBook(db, &book)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("could not create new book %v", err))
		}
		return c.JSON(book)
	}
}
