package models

import "github.com/gofiber/fiber/v2"

type QrInput struct {
	URLString string `json:"url" validate:"required,url"`
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
