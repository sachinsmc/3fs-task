package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	_ "github.com/sachinsmc/3fs-task/docs"
	"github.com/sachinsmc/3fs-task/middlewares"
)

// Setup
// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @contact.name API Support
// @contact.email hey@sachinsmc.me
// @host localhost:3003
// @BasePath /
func Setup(app *fiber.App) {

	middlewares.CORS(app)

	app.Get("/", monitor.New())

	app.Use(etag.New())
	app.Use(logger.New())

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")

	users := v1.Group("/users")

	users.Get("/", GetAll)
	users.Put("/:id", Update)
	users.Post("/", Create)
	users.Delete("/:id", Remove)

	groups := v1.Group("/groups")

	groups.Get("/", GetGroups)
	groups.Put("/:id", UpdateGroup)
	groups.Post("/", CreateGroup)
	groups.Delete("/:id", RemoveGroup)
}

type HTTPError struct {
	Status  string
	Message string
}
