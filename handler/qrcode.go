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

	png, errEncode := qrcode.Encode(inputUrl.URLString, qrcode.Highest, 1024)
	if errEncode != nil {
		defaultQrCode, _ := qrcode.Encode("https://www.default-url.com", qrcode.Highest, 1024)
		encoded := base64.StdEncoding.EncodeToString(defaultQrCode)
		return encoded
	} else {
		encoded := base64.StdEncoding.EncodeToString(png)
		return encoded
	}
}
