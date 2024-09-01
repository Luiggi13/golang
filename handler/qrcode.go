package handler

import (
	modelQR "goapi/models"
	u "goapi/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateQrCode(c *fiber.Ctx) interface{} {
	var inputUrl modelQR.QrInput

	if err := c.BodyParser(&inputUrl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(modelQR.ApiBadRequest)
	}

	if err := u.ValidateURL(inputUrl.URLString); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(modelQR.ApiBadRequest)
	}

	return u.GenerateQrCode(inputUrl.URLString)
}
