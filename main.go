package main

import (
	"errors"
	"goapi/router"
	utils "goapi/utils"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	_ = godotenv.Load(".env")
	listeningAddress := os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")
	// Fiber instance
	app := fiber.New(fiber.Config{
		AppName:      "Quick QR Code Generator " + os.Getenv("APIVERSION"),
		ServerHeader: "QQR",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			return ctx.Status(code).JSON(fiber.Map{
				"status": "[router] -> " + strconv.Itoa(code) + " " + err.Error(),
			})
		},
	})
	utils.Middlewares(app)
	// Setup App Routes
	router.CreateRoutes(app)

	// Start server
	log.Fatal(app.Listen(listeningAddress))
}
