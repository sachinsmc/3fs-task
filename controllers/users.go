package controllers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sachinsmc/3fs-task/db"
	"github.com/sachinsmc/3fs-task/models"
)

func InsertUser(u *models.Users) *models.Group {
	d := db.GetDB()
	group := &models.Group{
		ID: uuid.New(),
		Name: "name",
	}
	result := d.Create(group)
	if result.Error != nil {
		fmt.Println("Failed to create group : ", result.Error)
	}
	return group
}

func ListAllUsers() []models.Group {
	var groups []models.Group
	d := db.GetDB()
	d.Find(&groups)
	return groups
}



func UpdateUser(ID string) *models.Group {
	d := db.GetDB()
	group := &models.Group{
		ID: uuid.New(),
	}
	result := d.Create(group)
	if result.Error != nil {
		fmt.Println("Failed to create group : ", result.Error)
	}
	return group
}

func RemoveUser(name string) *Response {
	d := db.GetDB()
	group := &models.Group{
		ID: uuid.New(),
		Name: name,
	}
	result := d.Create(group)
	if result.Error != nil {
		fmt.Println("Failed to create group : ", result.Error)
	}
	return &Response{
		Message: "User deleted successfully.",
	}
}

