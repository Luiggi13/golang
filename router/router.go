package router

import (
	m "goapi/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// CreateRoutes for fiber app
func CreateRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Api v1
	v1 := api.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		err := c.Redirect("/api/v1/health", fiber.StatusMovedPermanently)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"status": "[v1] -> 301 Redirect"})
	})
	v1.Get("/health", health)
	v1.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	qrCodeRoutes := v1.Group("/qr", logger.New())
	qrCodeRoutes.Get("/", getAllQr)
	qrCodeRoutes.Get("/:id", getQrById)
	qrCodeRoutes.Post("/", createQr)
	qrCodeRoutes.Delete("/:id", deleteQrById)

}

// Endpoint Api v1
func health(c *fiber.Ctx) error {
	return c.JSON(m.GetHealth(c))
}
func createQr(c *fiber.Ctx) error {
	return c.JSON(m.CreateQrCode(c))
}
func getAllQr(c *fiber.Ctx) error {
	return c.JSON(m.GetAllQrCode(c))
}
func getQrById(c *fiber.Ctx) error {
	return c.JSON(m.GetByIdQrCode(c))
}

func deleteQrById(c *fiber.Ctx) error {
	return c.JSON(m.DeleteQrById(c))
}
