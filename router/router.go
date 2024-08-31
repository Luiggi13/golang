package router

import (
	m "goapi/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// CreateRoutes for fiber app
func CreateRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Api v1
	v1 := api.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		err := c.Redirect("/api/v1/info", fiber.StatusMovedPermanently)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"status": "[v1] -> 301 Redirect"})
	})
	v1.Get("/health", health)
	v1.Post("/qr/", createQr)

}

// Endpoint Api v1
func health(c *fiber.Ctx) error {
	return c.JSON(m.GetHealth(c))
}

func createQr(c *fiber.Ctx) error {
	return c.JSON(m.CreateQrCode(c))
}
