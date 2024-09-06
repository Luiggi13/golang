package handler

import (
	db "goapi/database"
	m "goapi/models"
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

func InitMigration(c *fiber.Ctx) interface{} {
	db.Connect_db(true, false, false)
	return m.MigrationInterface(m.BaseError{Message: "Migration deployed correctly", Method: c.Method()})
}
func InitSeeders(c *fiber.Ctx) interface{} {
	db.Connect_db(false, true, false)
	return m.MigrationInterface(m.BaseError{Message: "Seeder deployed correctly", Method: c.Method()})
}
func CleanTables(c *fiber.Ctx) interface{} {
	// db.Connect_db(false, true, true)
	return m.MigrationInterface(m.BaseError{Message: "Cleaned tables correctly", Method: c.Method()})
}
