package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/3fs-task/controllers"
	"github.com/sachinsmc/3fs-task/models"
)

// CreateGroup godoc
// @Summary Create group
// @Description add group
// @ID create group
// @Accept  json
// @Produce  json
// @Tags Group
// @Param group body models.CreateGroupRequest true "Group"
// @Success 200 {object} models.Group
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/groups [post]
func CreateGroup(c *fiber.Ctx) error {
	g := new(models.Group)
	if err := c.BodyParser(g); err != nil {
		fmt.Println("Failed to read body :", c.Body())
		return err
	}
	return c.JSON(controllers.InsertGroup(g.Name))
}

// GetGroups godoc
// @Summary Get all groups
// @Description Get group list
// @ID get-groups
// @Accept  json
// @Produce  json
// @Tags Group
// @Success 200 {object} models.Group
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/groups [get]
func GetGroups(c *fiber.Ctx) error {
	return c.JSON(controllers.ListAllGroups())
}

// UpdateGroup godoc
// @Summary update a group
// @Description Update Group by group id
// @ID update-group
// @Accept  json
// @Produce  json
// @Tags Group
// @Param id path string true "Group ID"
// @Param group body models.CreateGroupRequest true "Group"
// @Success 200 {object} models.Group
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/groups/{id} [put]
func UpdateGroup(c *fiber.Ctx) error {
	g := new(models.Group)
	if err := c.BodyParser(g); err != nil {
		fmt.Println("Failed to read body :", c.Body())
		return err
	}
	return c.JSON(controllers.UpdateGroup(*g, c.Params("id")))
}

// RemoveGroup godoc
// @Summary delete a group
// @Description Delete group by group id
// @ID delete-group
// @Accept  json
// @Produce  json
// @Tags Group
// @Param id path string true "Group ID"
// @Success 200 {object} controllers.Response
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/groups/{id} [delete]
func RemoveGroup(c *fiber.Ctx) error {
	return c.JSON(controllers.RemoveGroup(c.Params("id")))
}
