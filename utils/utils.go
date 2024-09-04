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

// ValidateURL checks if a given URL string is valid.
//
// The function takes a single parameter:
//   - urlString: A string representing the URL to be validated.
//
// The function returns an error if the URL is invalid.
// If the URL is valid, it returns nil.
//
// The function uses the url.ParseRequestURI function to parse the URL string.
// It checks if the parsing error is not nil, or if the URL scheme or host are empty.
// If any of these conditions are met, the function returns an error with the message "invalid URL format".
// Otherwise, it returns nil.
// Commit that removes ValidateURL function: 2763761832f4f93c3275c360cf81c43affeedb33
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
		return base64.StdEncoding.EncodeToString(defaultQrCode)
	}

	return base64.StdEncoding.EncodeToString(png)
}

// Middlewares sets up and applies middleware functions to the Fiber app.
// It includes CORS and ETag middleware configurations.
//
// app: A pointer to the Fiber app instance.
func Middlewares(app *fiber.App) {
	// Apply CORS middleware with the specified configuration.
	// This allows cross-origin requests from any origin.
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	// Apply ETag middleware with default configuration.
	// This generates an ETag header for each response.
	app.Use(etag.New())

	// Apply ETag middleware with custom configuration.
	// This generates a weak ETag header for each response.
	app.Use(etag.New(etag.Config{
		Weak: true,
	}))
}

// IsUserProvided checks if the user ID is provided in the QR input based on the specified method.
//
// The function takes two parameters:
//   - user: A QrInput struct representing the user input for generating a QR code.
//     It contains a UserId field which is a pointer to a string.
//   - method: A string representing the method used to generate the QR code.
//
// The function returns a boolean value indicating whether the user ID is provided.
// If the UserId field in the user struct is nil or its value is an empty string,
// the function returns false. Otherwise, it returns true.
func IsUserProvided(user m.QrInput, method string) bool {
	if user.UserId == nil || *user.UserId == "" {
		return false
	}
	return true
}
