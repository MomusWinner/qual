package middleware

import (
	"app/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func MetricsMiddleware(m domain.HttpMetrics) fiber.Handler {
	return func(c *fiber.Ctx) error {
		statusCode := c.Response().StatusCode()
		method := c.Method()
		path := c.Path()

		m.StartRequestMetrics(statusCode, method, path)
		err := c.Next()
		m.EndRequestMetrics()

		return err
	}
}
