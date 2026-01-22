package main

import (
	"app/internal/api/v1/user"
	"app/internal/core"
	"app/internal/domain/cases"
	"fmt"

	// flogger "app/internal/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Service API
// @version 1.0
// @description Service api :-)
// @basePath /api/v1
func main() {
	ctx := core.InitCtx()
	userUseCase := cases.NewUserUseCase(ctx)
	userHandler := user.NewUserHandler(ctx, userUseCase)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language," +
			"Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "http://localhost:3000,http://localhost:8000,https://*.ocrv-game.ru,https://ocrv-game.ru",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	swaggerCfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}

	app.Use(swagger.New(swaggerCfg))

	api := fiber.New()
	app.Mount("/api/v1", api)

	user.AddRoutes(api, userHandler)

	app.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

	app.Listen(ctx.Config().GetHost())
	// userMiddleware := middleware.NewUserMiddleware(user_repository, conf)

	//
	// app := fiber.New()
	// micro := fiber.New()
	// app.Use(cors.New(cors.Config{
	// 	AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language," +
	// 		"Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
	// 	AllowOrigins:     "http://localhost:3000,http://localhost:8000,https://*.ocrv-game.ru,https://ocrv-game.ru",
	// 	AllowCredentials: true,
	// 	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	// }))
	//
	// swagger_conf := swagger.Config{
	// 	BasePath: "/",
	// 	FilePath: "./docs/swagger.json",
	// 	Path:     "swagger",
	// 	Title:    "Swagger API Docs",
	// }
	//
	// app.Use(swagger.New(swagger_conf))
	// loggerMiddelware := flogger.NewLoggerMiddelware(conf.Debug)
	// app.Use(loggerMiddelware.Handle)
	//
	// app.Mount("/api/v1", micro)
	// // userMiddleware := middleware.NewUserMiddleware(user_repository, conf)
	//
	// authHandler := auth.NewAuthHandler(user_repository, &redis_conn, &smtp, conf)
	// auth.AddRoutes(micro, &authHandler)
	//
	// convImpl := &DTO.UserConverterImpl{}
	// userHandler := user.NewUserHandler(conf, &smtp, user_repository, &redis_conn, convImpl)
	// user.AddRoutes(micro, userHandler)
	//
	// micro.Get("/", func(c *fiber.Ctx) error {
	// 	slog.InfoContext(c.Context(), "Hellooooo")
	// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 		"status":  "success",
	// 		"message": "Hello",
	// 	})
	// })
	//
	// micro.Get("/healthcheck", func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 		"status":  "success",
	// 		"message": "good",
	// 	})
	// })
	//
	// micro.All("*", func(c *fiber.Ctx) error {
	// 	path := c.Path()
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"status":  "fail",
	// 		"message": fmt.Sprintf("Path: %v does not exists on this server", path),
	// 	})
	// })
	//
	// app.Listen(conf.Host)
}
