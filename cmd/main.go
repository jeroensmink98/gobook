package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jeroensmink98/gobook/pkg/db"
	"github.com/jeroensmink98/gobook/pkg/handlers"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database := db.ConnectDB()

	db.Delete((database))
	db.InitDb((database))
	db.SeedDb((database))

	app.Get("/api/book", handlers.GetBooks((database)))
	app.Get("/api/book/:id", handlers.GetBookById((database)))
	app.Post("/api/book", handlers.CreateBook((database)))

	app.Get("/api/genre", handlers.GetGenres((database)))
	app.Get("/api/genre/:id", handlers.GetGenreById((database)))
	app.Post("/api/genre", handlers.CreateGenre((database)))

	log.Fatal(app.Listen(("localhost:8080")))
}
