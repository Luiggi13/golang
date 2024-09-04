package models

import "github.com/gofiber/fiber/v2"

type QrInput struct {
	URLString string  `json:"url" validate:"required,url"`
	UserId    *string `json:"userId"`
	Premium   *bool   `json:"premium"`
}

type RequestMethod string

type ApiError struct {
	Method RequestMethod `json:"method"`
	Error  string        `json:"error"`
}

var ApiRecordNotFound = ApiError{Method: "", Error: "404 Record not found"}
var ApiBadRequest = ApiError{Method: "", Error: "405 Bad Request"}
var ApiInternalError = ApiError{Method: "", Error: "500 Internal Server Error"}

func (m *RequestMethod) SetMethod(c *fiber.Ctx, p ...string) {
	*m = RequestMethod(SetMethod(c, p...))
}

func SetMethod(c *fiber.Ctx, p ...string) string {
	var ApiMessagePrefix = "[v1] -> "
	if p == nil {
		p = []string{""}
	}
	switch p[0] {
	case "JSON":
		return ApiMessagePrefix + c.Method() + " with " + "JSON"
	case "Query":
		return ApiMessagePrefix + c.Method() + " with " + "Query Parameter"
	case "Path":
		return ApiMessagePrefix + c.Method() + " with " + "Path Parameter"
	default:
		return ApiMessagePrefix + c.Method()
	}
}

type QrCodeGenerated struct {
	Id         string `json:"id"`
	StatusCode int64  `json:"status_code"`
	QrCode     string `json:"qr_code"`
	Premium    bool   `json:"premium" validate:"required,premium"`
}

type BaseError struct {
	Method  string `json:"method"`
	Message string `json:"message"`
	Url     string `json:"url,omitempty"`
}

type CustomErrorQR CustomErrorQRElement

type CustomErrorQRElement struct {
	Status     string  `json:"status"`
	StatusCode int64   `json:"status_code"`
	Message    string  `json:"message"`
	Details    Details `json:"details"`
}

type Details struct {
	Method string `json:"method"`
	URL    string `json:"url,omitempty"`
}

func BadRequestError(ue BaseError) CustomErrorQR {

	details := Details{
		Method: ue.Method,
	}
	if ue.Url != "" {
		details.URL = ue.Url
	}
	return CustomErrorQR{
		Status:     "Bad Request",
		StatusCode: 400,
		Message:    ue.Message,
		Details:    details,
	}
}
