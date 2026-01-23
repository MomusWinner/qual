package api

import (
	"app/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func getAppContext(c *fiber.Ctx) domain.Context {
	if ctx, ok := c.Locals("app_context").(domain.Context); ok {
		return ctx
	}
	return nil
}

func ContextWrapper(handler func(domain.Context, *fiber.Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := getAppContext(c)
		if ctx == nil {
			panic("domain.Context is nil")
		}
		return handler(ctx, c)
	}
}
