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

	validURL := ValidateURL(DEFAULTURL)
	assert.Equal(t, nil, validURL)

	invalidText := strings.Split(DEFAULTURL, "://")
	invalidURL := ValidateURL(invalidText[1])
	assert.Equal(t, "invalid URL", invalidURL.Error())

	app.ReleaseCtx(c)
}
