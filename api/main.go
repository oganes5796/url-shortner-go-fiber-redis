package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/oganes5796/url-shorten-fiber-redis/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		// Send a string  to the client
		return c.SendString("Hello, World !")
	})
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Info("Error loading .env file", "err", err)
	}

	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
