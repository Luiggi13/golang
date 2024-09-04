package handler

import (
	m "goapi/models"
	u "goapi/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateQrCode(c *fiber.Ctx) interface{} {
	var inputUrl m.QrInput
	var isPremium = false

	if err := c.BodyParser(&inputUrl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(m.BaseError{Message: "405 Bad Request", Method: c.Method()})
	}

	if !u.IsUserProvided(inputUrl, c.Method()) {
		return m.BadRequestError(m.BaseError{Message: "Error user", Method: c.Method()})
	}

	if err := u.ValidateURL(inputUrl.URLString); err != nil {
		return m.BadRequestError(m.BaseError{Message: "Error user", Method: c.Method(), Url: inputUrl.URLString})
	}

	if inputUrl.Premium != nil {
		isPremium = *inputUrl.Premium
	}

	response := m.QrCodeGenerated{
		Id:         *inputUrl.UserId,
		StatusCode: 200,
		QrCode:     u.GenerateQrCode(inputUrl.URLString),
		Premium:    isPremium,
	}
	return response
}
