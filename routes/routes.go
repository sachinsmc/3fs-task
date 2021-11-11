package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/sachinsmc/3fs-task/middlewares"
)


func Setup(app *fiber.App) {

	middlewares.CORS(app)

	app.Get("/", monitor.New())

	app.Use(etag.New())
	app.Use(logger.New())

	api := app.Group("/api")

	v1 := api.Group("/v1")

	users := v1.Group("/users")

	users.Get("/", GetAll)
	users.Put("/", Update)
	users.Post("/", Create)
	users.Delete("/", Remove)

	groups := v1.Group("/groups")

	groups.Get("/", GetGroups)
	groups.Put("/:id", UpdateGroup)
	groups.Post("/", CreateGroup)
	groups.Delete("/:id", RemoveGroup)
}
