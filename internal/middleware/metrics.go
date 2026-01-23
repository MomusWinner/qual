package middleware

import (
	"app/internal/core"

	"github.com/gofiber/fiber/v2"
)

func MetricsMiddleware(m *core.HttpMetrics) fiber.Handler {
	return func(c *fiber.Ctx) error {
		statusCode := c.Response().StatusCode()
		method := c.Method()
		path := c.Path()

		r := m.StartRequestMetrics(statusCode, method, path)
		err := c.Next()
		r.EndRequestMetrics()

		return err
	}
}
