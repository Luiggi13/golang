package handler

import (
	"encoding/base64"
	"fmt"
	modelQR "goapi/models"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

func validateURL(urlString string) error {
	u, err := url.ParseRequestURI(urlString)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("invalid URL")
	}
	return nil
}

func CreateQrCode(c *fiber.Ctx) interface{} {
	var inputUrl modelQR.QrInput

	if err := c.BodyParser(&inputUrl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(modelQR.ApiBadRequest)
	}

	if err := validateURL(inputUrl.URLString); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(modelQR.ApiBadRequest)
	}

	// Generar el código QR
	png, err3 := qrcode.Encode(inputUrl.URLString, qrcode.Highest, 1024)
	if err3 != nil {
		qr, _ := qrcode.Encode(inputUrl.URLString, qrcode.Highest, 1024)
		encoded := base64.StdEncoding.EncodeToString(qr)
		return encoded
	} else {
		encoded := base64.StdEncoding.EncodeToString(png)
		return encoded
	}
}
