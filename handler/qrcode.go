package handler

import (
	db "goapi/database"
	m "goapi/models"
	u "goapi/utils"
	"strconv"

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
		return c.Status(fiber.StatusBadRequest).JSON(m.BaseError{
			Message: "405 Bad Request",
			Method:  c.Method(),
		})
	}

	if !u.IsUserProvided(inputUrl, c.Method()) {
		return c.Status(fiber.StatusBadRequest).JSON(m.BaseError{
			Message: "Error user",
			Method:  c.Method(),
		})
	}

	if inputUrl.Premium != nil {
		isPremium = *inputUrl.Premium
	}

	myInsert := m.QRStruct{
		QrCode:  u.GenerateQrCode(inputUrl.URLString),
		UserId:  inputUrl.UserId,
		UrlText: inputUrl.URLString,
		Premium: isPremium,
		IdTag:   inputUrl.IdTag,
	}

	err := db.InsertQR(myInsert)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(m.BaseError{
			Message: "Error inserting QR code",
			Method:  c.Method(),
		})
	}

	response := m.QrCodeGenerated{
		Id:         *inputUrl.UserId,
		StatusCode: 200,
		QrCode:     myInsert.QrCode,
		Premium:    isPremium,
		TagName:    "",
	}

	rows := getTagForPost(c, inputUrl.IdTag)
	defer rows.Close()

	var qr m.SelectTags
	if rows.Next() {

		err := rows.Scan(&qr.TagId, &qr.TagName, &qr.Public)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(m.BaseError{
				Message: "Error scanning tag result",
				Method:  c.Method(),
			})
		}
		response.TagName = qr.TagName
	}

	return response
}

func GetAllQrCode(c *fiber.Ctx) []m.QRStructJoin {
	var qrList []m.QRStructJoin = []m.QRStructJoin{}
	rows := db.GetAll(c)

	for rows.Next() {
		var qr m.QRStructJoin
		rows.Scan(&qr.QrId, &qr.QrCode, &qr.UserId, &qr.UrlText, &qr.Premium, &qr.TagName)
		qrList = append(qrList, qr)
	}

	return qrList
}

func GetByIdQrCode(c *fiber.Ctx) interface{} {
	var qr m.QRStructJoin

	rows := db.GetById(c)
	for rows.Next() {
		err := rows.Scan(&qr.QrId, &qr.QrCode, &qr.UserId, &qr.UrlText, &qr.Premium, &qr.TagName)

		if err == nil {
			return qr
		}
	}

	return m.NotFound(m.BaseError{Message: "Non QR code", Method: c.Method()})
}

func DeleteQrById(c *fiber.Ctx) interface{} {
	_, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return m.BadRequestError(m.BaseError{Message: "Id param should be a number", Method: c.Method()})
	}
	res, err := db.DeleteById(c)
	if err != nil {
		return m.InternalRequestError(m.BaseError{Message: "Internal server error. Try again in a few minutes", Method: c.Method()})
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return m.InternalRequestError(m.BaseError{Message: "Internal server error. Try again in a few minutes", Method: c.Method()})
	}

	switch rowsAffected {
	case 0:
		return m.NotFound(m.BaseError{Message: "Failed to delete QR code", Method: c.Method()})
	default:
		return m.DeleteResponse(m.BaseError{Message: "Resource deleted", Method: c.Method()})
	}
}
