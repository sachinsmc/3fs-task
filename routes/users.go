package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/3fs-task/controllers"
	"github.com/sachinsmc/3fs-task/models"
)

// Create godoc
// @Summary Create user
// @Description add user
// @ID get-users
// @Accept  json
// @Produce  json
// @Tags User
// @Param user body models.CreateUserRequest true "User"
// @Success 200 {object} models.Users
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/users [post]
func Create(c *fiber.Ctx) error {
	u := new(models.Users)
	if err := c.BodyParser(u); err != nil {
		fmt.Println("Failed to read body :", c.Body())
		return err
	}
	return c.JSON(controllers.InsertUser(u))
}

// GetAll godoc
// @Summary Get all users
// @Description Get users list
// @ID get-users
// @Accept  json
// @Produce  json
// @Tags User
// @Success 200 {object} models.Users
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/users [get]
func GetAll(c *fiber.Ctx) error {
	return c.JSON(controllers.ListAllUsers())
}

// Update godoc
// @Summary update a user
// @Description Update user by user id
// @ID update-user
// @Accept  json
// @Produce  json
// @Tags User
// @Param id path string true "User ID"
// @Param user body models.CreateUserRequest true "User"
// @Success 200 {object} models.Users
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/users/{id} [put]
func Update(c *fiber.Ctx) error {
	u := new(models.Users)
	if err := c.BodyParser(u); err != nil {
		fmt.Println("Failed to read body :", c.Body())
		return err
	}
	return c.JSON(controllers.UpdateUser(*u, c.Params("id")))
}

// Remove godoc
// @Summary delete a user
// @Description Delete user by user id
// @ID delete-user
// @Accept  json
// @Produce  json
// @Tags User
// @Param id path string true "User ID"
// @Success 200 {object} controllers.Response
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/users/{id} [delete]
func Remove(c *fiber.Ctx) error {
	return c.JSON(controllers.RemoveUser(c.Params("id")))
}
