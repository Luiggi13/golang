package handler

import (
	db "goapi/database"
	m "goapi/models"
	u "goapi/utils"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" // add this
)

// CreateQrCode is a handler function that generates a QR code based on the provided URL.
// It accepts a Fiber context (c *fiber.Ctx) and returns an interface{}.
//
// The function performs the following steps:
// 1. Decodes the request body into a QrInput struct (var inputUrl m.QrInput).
// 2. Checks if the request body contains a premium field and sets the isPremium variable accordingly.
// 4. Checks if the user is authenticated using the IsUserProvided function from the utils package.
// 5. If any validation fails, it returns a 400 Bad Request response with an appropriate error message.
// 6. If the user is authenticated and the URL is valid, it generates a QR code using the GenerateQrCode function from the utils package.
// 7. Constructs a QrCodeGenerated struct with the user ID, status code, generated QR code, and premium status.
// 8. Returns the constructed QrCodeGenerated struct as the response.

func CreateQrCode(c *fiber.Ctx) interface{} {
	var inputUrl m.QrInput
	var isPremium = false

	if err := c.BodyParser(&inputUrl); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(m.BaseError{Message: "405 Bad Request", Method: c.Method()})
	}

	if !u.IsUserProvided(inputUrl, c.Method()) {
		return m.BadRequestError(m.BaseError{Message: "Error user", Method: c.Method()})
	}

	if inputUrl.Premium != nil {
		isPremium = *inputUrl.Premium
	}

	myInsert := m.QRStruct{
		QrCode:  u.GenerateQrCode(inputUrl.URLString),
		User:    inputUrl.UserId,
		Premium: isPremium,
		Url:     inputUrl.URLString,
	}
	response := m.QrCodeGenerated{
		Id:         *inputUrl.UserId,
		StatusCode: 200,
		QrCode:     u.GenerateQrCode(inputUrl.URLString),
		Premium:    isPremium,
	}
	db.InsertQR(myInsert)
	return response
}

func GetAllQrCode(c *fiber.Ctx) []m.QRStruct {
	var qrList []m.QRStruct = []m.QRStruct{}
	rows := db.GetAll(c)

	for rows.Next() {
		var qr m.QRStruct

		rows.Scan(&qr.User, &qr.Premium, &qr.QrCode)
		qrList = append(qrList, qr)
	}

	return qrList
}

func GetByIdQrCode(c *fiber.Ctx) interface{} {
	var qr m.QRStruct

	rows := db.GetById(c)
	for rows.Next() {
		err := rows.Scan(&qr.User, &qr.QrCode, &qr.Premium)

		if err == nil {
			return qr
		}
	}

	return m.NotFound(m.BaseError{Message: "Non-existent user", Method: c.Method()})
}
