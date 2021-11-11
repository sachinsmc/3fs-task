package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/3fs-task/controllers"
	"github.com/sachinsmc/3fs-task/models"
)

func Create(c *fiber.Ctx) error {
	u := new(models.Users)
	if err := c.BodyParser(u); err != nil {
		fmt.Println("Failed to read body :", c.Body())
		return err
	}
	return c.JSON(controllers.InsertUser(u))
}

func GetAll(c *fiber.Ctx) error {
	return c.JSON(controllers.ListAllUsers())
}

func Update(c *fiber.Ctx) error {
	return c.JSON(controllers.UpdateUser(c.Params("id")))
}

func Remove(c *fiber.Ctx) error {
	return c.JSON(controllers.UpdateUser(c.Params("id")))
}

