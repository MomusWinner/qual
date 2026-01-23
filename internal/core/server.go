package core

import (
	"app/internal/api/v1/user"
	"app/internal/domain"
	"app/internal/domain/cases"
	"app/internal/middleware"
	"fmt"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	ctx domain.Context
	app *fiber.App
}

func NewServer(ctx domain.Context, enableMetrics bool, enableSwagger bool) *Server {
	if enableMetrics {
		ctx.Connection().EnableUserRepositoryMetrics()
	}
	userUseCase := cases.NewUserUseCase(ctx)
	userHandler := user.NewUserHandler(userUseCase)

	app := fiber.New()

	if enableMetrics {
		RegisterMetricsAt(app, "/metrics", middleware.CorrelationIDMiddleware(ctx))
	}

	if enableSwagger {
		swaggerCfg := swagger.Config{
			BasePath: "/",
			FilePath: "./docs/swagger.json",
			Path:     "swagger",
			Title:    "Swagger API Docs",
		}

		app.Use(swagger.New(swaggerCfg))
	}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
	})

	api := fiber.New()
	app.Mount("/api/v1", api)
	if enableMetrics {
		api.Use(middleware.MetricsMiddleware(ctx.HttpMetrics()))
	}
	api.Use(middleware.CorrelationIDMiddleware(ctx))
	user.AddRoutes(api, userHandler)

	app.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

	return &Server{
		ctx: ctx,
		app: app,
	}
}

func (s *Server) Start() {
	s.app.Listen(s.ctx.Config().GetHost())
}

// ctx := core.InitCtx()
// ctx.Connection().EnableUserRepositoryMetrics()
// userUseCase := cases.NewUserUseCase(ctx)
// userHandler := user.NewUserHandler(userUseCase)
//
// metrics := core.NewHttpMetrics()
//
// app := fiber.New()
//
// app.Use(cors.New(cors.Config{
// 	AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language," +
// 		"Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
// 	AllowOrigins:     "http://localhost:3000,http://localhost:8000,https://*.ocrv-game.ru,https://ocrv-game.ru",
// 	AllowCredentials: true,
// 	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
// }))
//
// core.RegisterMetricsAt(app, "/metrics", middleware.CorrelationIDMiddleware(ctx))
//
// swaggerCfg := swagger.Config{
// 	BasePath: "/",
// 	FilePath: "./docs/swagger.json",
// 	Path:     "swagger",
// 	Title:    "Swagger API Docs",
// }
//
// app.Get("/healthcheck", func(c *fiber.Ctx) error {
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
// })
//
// app.Use(swagger.New(swaggerCfg))
//
// api := fiber.New()
// app.Mount("/api/v1", api)
// api.Use(middleware.MetricsMiddleware(metrics))
// api.Use(middleware.CorrelationIDMiddleware(ctx))
// user.AddRoutes(api, userHandler)
//
// app.All("*", func(c *fiber.Ctx) error {
// 	path := c.Path()
// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 		"status":  "fail",
// 		"message": fmt.Sprintf("Path: %v does not exists on this server", path),
// 	})
// })
//
// app.Listen(ctx.Config().GetHost())
