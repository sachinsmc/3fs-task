package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/3fs-task/config"
	"github.com/sachinsmc/3fs-task/db"
	"github.com/spf13/viper"
)

func Run() {
	config.Init()

	db.ConnectDB()

	app := fiber.New()

	app.Listen(":" + viper.GetString("server.port"))
}
