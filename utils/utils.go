package utils

import (
	"encoding/base64"
	"fmt"
	m "goapi/models"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/skip2/go-qrcode"
)

const DefaultUrl = "https://www.default-url.com"

func ValidateURL(urlString string) error {
	u, err := url.ParseRequestURI(urlString)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("invalid URL format")
	}
	return nil
}

func GenerateQrCode(url string) string {
	png, errEncode := qrcode.Encode(url, qrcode.Highest, 1024)
	if errEncode != nil {
		defaultQrCode, _ := qrcode.Encode(DefaultUrl, qrcode.Highest, 1024)
		encoded := base64.StdEncoding.EncodeToString(defaultQrCode)
		return encoded
	} else {
		encoded := base64.StdEncoding.EncodeToString(png)
		return encoded
	}
}

func Middlewares(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	app.Use(etag.New())
	app.Use(etag.New(etag.Config{
		Weak: true,
	}))
}

func IsUserProvided(user m.QrInput, method string) bool {
	if user.UserId == nil || *user.UserId == "" {
		return false
	}
	return true
}
