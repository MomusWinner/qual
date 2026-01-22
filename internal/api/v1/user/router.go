package user

import (
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(router fiber.Router, handler *UserHandler) {
	user := router.Group("/users")

	user.Post("/", handler.Create)
	user.Get("/:id", handler.GetById)
	user.Get("/", handler.GetAll)
	user.Put("/:id", handler.Update)
	user.Delete("/:id", handler.Delete)
}
