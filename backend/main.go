package main

import (
	"fmt"
	db "github.com/amalshaji/stoyle.me/database"
	"github.com/amalshaji/stoyle.me/handlers"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func setupMiddlewares(app *fiber.App) {
	app.Use(csrf.New(csrf.Config{
		Expiration: 1*time.Hour,
	}))
	app.Use(logger.New())
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Post("/api/checkavailability", handlers.CheckAvailability)
	app.Post("/api/shorten", handlers.Shorten)
	app.Get("/:short", handlers.ResolveShort)
}

func initDB() {
	db.DB = redis.NewClient(&redis.Options{
		Addr:               os.Getenv("REDIS_ADDRESS"),
		Password:           os.Getenv("REDIS_PASSWORD"),
		DB:                 0,
	})
	err := db.DB.Ping(db.Ctx).Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Started Db connection")
}

func loadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	engine := html.New("../frontend/views", ".html")

	initDB()
	defer db.DB.Close()

	loadEnvs()

	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/", "../frontend/public")

	setupMiddlewares(app)
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
