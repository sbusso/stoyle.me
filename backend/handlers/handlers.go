package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CheckAvailability to check the availability of custom short
func CheckAvailability(c *fiber.Ctx) error {
	v := rand.Intn(100)
	time.Sleep(1 * time.Second)
	return c.Status(200).JSON(fiber.Map{
		"message": v%2 == 0,
	})
}


// Shorten creates short links
func Shorten(c *fiber.Ctx) error {
	time.Sleep(1 * time.Second)
	id := fmt.Sprintf("%s", uuid.New())[:6]
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"short": id,
	})
}
