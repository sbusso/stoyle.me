package handlers

import (
	db "github.com/amalshaji/stoyle.me/database"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"

	"github.com/gofiber/fiber/v2"
)

type availRequest struct {
	Short string `json:"short"`
}

// CheckAvailability to check the availability of custom short
func CheckAvailability(c *fiber.Ctx) error {
	var req availRequest

	err := c.BodyParser(&req)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Bad request")
	}

	val, err := db.DB.Get(db.Ctx, req.Short).Result()
	if val == "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": false,
	})
}

type request struct {
	URL string `json:"url"`
	CustomShort string `json:"short"`
	Expiry	int `json:"expires"`
}

// Shorten creates short links
func Shorten(c *fiber.Ctx) error {
	var req request
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Wrong body format")
	}

	if req.CustomShort == "" {
		req.CustomShort = uuid.New().String()[:6]
	}

	// Check if the short is present in DB
	val, _ := db.DB.Get(db.Ctx, req.CustomShort).Result()


	if val != "" {
		return c.Status(fiber.StatusBadRequest).SendString("Short URL already in use")
	}

	err = db.DB.Set(db.Ctx, req.CustomShort, req.URL, time.Duration(req.Expiry)*time.Hour).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error connection to Database")
	}

	return c.Status(fiber.StatusOK).JSON(req)
}

// ResolveShort redirects to original URL
func ResolveShort(c *fiber.Ctx) error {
	short := c.Params("short")

	v, err := db.DB.Get(db.Ctx, short).Result()
	if err == redis.Nil {
		return c.Redirect("https://stoyle.me", 302)
	}
	return c.Redirect(v, 302)
}
