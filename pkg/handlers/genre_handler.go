package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jeroensmink98/gobook/pkg/dbops"
	"github.com/jeroensmink98/gobook/pkg/models"
	"gorm.io/gorm"
)

func GetGenres(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		genres, err := dbops.GetAllGenres(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error retrieving genres: %v", err))

		}
		return c.JSON(genres)
	}
}

func GetGenreById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("invalid id")
		}

		genre, err := dbops.GetGenreByID(db, uint(id))
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("genre with id %v not found", id))
		}
		return c.JSON(genre)
	}
}

func CreateGenre(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var genre models.Genre

		if err := c.BodyParser(&genre); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("parse error: %v", err))
		}

		err := dbops.CreateGenre(db, &genre)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("could not create new genre %v", err))
		}
		return c.JSON(genre)
	}
}
