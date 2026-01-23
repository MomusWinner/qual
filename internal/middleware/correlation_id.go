package middleware

import (
	"app/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CorrelationIDMiddleware(appContext domain.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		correlationID := c.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String()
		}
		c.Set("X-Correlation-ID", correlationID)
		appCtx := appContext.WithCorrelationID(correlationID)

		c.Locals("app_context", appCtx)

		appCtx.Logger().Info("request started",
			"method", c.Method(),
			"path", c.Path(),
			"ip", c.IP(),
		)

		err := c.Next()

		return err
	}
}
