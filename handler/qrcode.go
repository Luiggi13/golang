package handler

import (
	m "goapi/models"
	u "goapi/utils"

	"github.com/gofiber/fiber/v2"
)

type ErrorValidate2 struct {
	Status     string            `json:"status"`
	StatusCode int               `json:"status_code"`
	Message    string            `json:"message"`
	Details    map[string]string `json:"details,omitempty"`
}

func CreateQrCode(c *fiber.Ctx) interface{} {
	var inputUrl m.QrInput

	if err := c.BodyParser(&inputUrl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(m.ApiBadRequest)
	}

	if err := u.ValidateURL(inputUrl.URLString); err != nil {
		return m.ErrorBadRequest(c.Method(), inputUrl.URLString)
	}

	return u.GenerateQrCode(inputUrl.URLString)
}
