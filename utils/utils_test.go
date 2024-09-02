package utils

import (
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestValidateURL(t *testing.T) {
	app := fiber.New()
	fctx := new(fasthttp.RequestCtx)
	c := app.AcquireCtx(fctx)

	validURL := ValidateURL(DefaultUrl)
	assert.Equal(t, nil, validURL)

	invalidText := strings.Split(DefaultUrl, "://")
	invalidURL := ValidateURL(invalidText[1])
	assert.Equal(t, "invalid URL format", invalidURL.Error())

	app.ReleaseCtx(c)
}
