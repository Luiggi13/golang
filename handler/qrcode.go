package handler

import (
	"encoding/base64"
	modelQR "goapi/models"
	u "goapi/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

func CreateQrCode(c *fiber.Ctx) interface{} {
	var inputUrl modelQR.QrInput

	if err := c.BodyParser(&inputUrl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(modelQR.ApiBadRequest)
	}

	if err := u.ValidateURL(inputUrl.URLString); err != nil {
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
