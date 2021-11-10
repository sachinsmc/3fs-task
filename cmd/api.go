package cmd

import (
	"github.com/gofiber/fiber/v2"
)

func Run()  {
	app := fiber.New()
	app.Listen("")
}
