package main

import (
	"github.com/amalshaji/stoyle.me/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func setupMiddlewares(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{AllowOrigins: "http://localhost:5000"}))
	app.Use(csrf.New())
}

func setupRoutes(app *fiber.App) {
	app.Post("/api/checkavailability", handlers.CheckAvailability)
	app.Post("/api/shorten", handlers.Shorten)
}

func main() {
	app := fiber.New()

	app.Static("/", "../frontend/public")

	setupMiddlewares(app)
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
