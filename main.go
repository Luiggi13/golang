package main

import (
	"errors"
	"goapi/router"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Load .env file
	_ = godotenv.Load(".env")
	listeningAddress := os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")

	// Fiber instance
	app := fiber.New(fiber.Config{
		AppName:        "Apir to generetae QR for urls",
		ServerHeader:   "QR Code 313",
		RequestMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
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
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Use(etag.New())
	app.Use(etag.New(etag.Config{
		Weak: true,
	}))
	// Setup App Routes
	router.CreateRoutes(app)

	// Start server
	log.Fatal(app.Listen(listeningAddress))
}
