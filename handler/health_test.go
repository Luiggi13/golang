package handler

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestGetHealth(t *testing.T) {
	// Create a Fiber instance
	app := fiber.New()

	// Create a mock Fasthttp request context
	fctx := new(fasthttp.RequestCtx)

	// Acquire a Fiber context from the Fasthttp context
	c := app.AcquireCtx(fctx)

	// Ejecutar la función GetHealth
	health := GetHealth(c)

	// Verify that the returned status is 200
	assert.Equal(t, 200, health.Status)

	// Release the Fiber context to avoid memory leaks
	app.ReleaseCtx(c)
}
