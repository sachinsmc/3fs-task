package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/3fs-task/controllers"
	"github.com/sachinsmc/3fs-task/models"
)

func CreateGroup(c *fiber.Ctx) error {
	g := new(models.Group)
	if err := c.BodyParser(g); err != nil {
		fmt.Println("Failed to read body :", c.Body())
		return err
	}
	return c.JSON(controllers.InsertGroup(g.Name))
}

func GetGroups(c *fiber.Ctx) error {
	return c.JSON(controllers.ListAllGroups())
}

func UpdateGroup(c *fiber.Ctx) error {
	g := new(models.Group)
	if err := c.BodyParser(g); err != nil {
		fmt.Println("Failed to read body :", c.Body())
		return err
	}
	return c.JSON(controllers.UpdateGroup(*g, c.Params("id")))
}

func RemoveGroup(c *fiber.Ctx) error {
	return c.JSON(controllers.RemoveGroup(c.Params("id")))
}