package handler

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type Health struct {
	Status     int    `json:"status"`
	ApiVersion string `json:"version"`
}

func GetHealth(c *fiber.Ctx) Health {
	return Health{
		Status:     200,
		ApiVersion: os.Getenv("APIVERSION"),
	}
}
