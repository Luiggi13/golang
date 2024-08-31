package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Health struct {
	Status int `json:"status"`
}

func GetHealth(c *fiber.Ctx) Health {
	return Health{
		Status: 200,
	}
}
