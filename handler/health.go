package handler

import (
	db "goapi/database"
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

func InitMigration(c *fiber.Ctx) error {
	db.Connect_db(true, false)
	return nil
}
func InitSeeders(c *fiber.Ctx) error {
	db.Connect_db(false, true)
	return nil
}
