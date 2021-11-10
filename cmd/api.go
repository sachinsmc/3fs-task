package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/3fs-task/config"
	"github.com/spf13/viper"
)

func Run() {
	config.Init()
	app := fiber.New()
	app.Listen(":"+viper.GetString("server.port"))
}
