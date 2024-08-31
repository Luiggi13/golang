package handler

import (
	"encoding/base64"
	"fmt"
	modelQR "goapi/models"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

func CreateQrCode(c *fiber.Ctx) interface{} {
	var inputUrl modelQR.QrInput

	if err := c.BodyParser(&inputUrl); err != nil {
		c.Status(fiber.StatusBadRequest)
		modelQR.ApiBadRequest.Method.SetMethod(c)
		return modelQR.ApiBadRequest
	}

	u, err2 := url.ParseRequestURI(inputUrl.URLString)
	if err2 != nil || u.Scheme == "" || u.Host == "" {
		fmt.Println("hay error")
		fmt.Println(u)
		c.Status(fiber.StatusBadRequest)
		modelQR.ApiBadRequest.Method.SetMethod(c)
		return modelQR.ApiBadRequest
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
