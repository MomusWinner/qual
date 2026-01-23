package user

import (
	"app/internal/api"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(router fiber.Router, handler *UserHandler) {
	user := router.Group("/users")

	user.Post("/", api.ContextWrapper(handler.Create))
	user.Get("/:id", api.ContextWrapper(handler.GetById))
	user.Get("/", api.ContextWrapper(handler.GetAll))
	user.Put("/:id", api.ContextWrapper(handler.Update))
	user.Delete("/:id", api.ContextWrapper(handler.Delete))
}
