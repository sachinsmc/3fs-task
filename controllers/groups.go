package controllers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sachinsmc/3fs-task/db"
	"github.com/sachinsmc/3fs-task/models"
)

//type Data struct {
//	result interface{}
//}

type Response struct {
	Message string `json:"message"`
}

func InsertGroup(name string) *models.Group {
	d := db.GetDB()
	group := &models.Group{
		ID:   uuid.New(),
		Name: name,
	}
	result := d.Create(group)
	if result.Error != nil {
		fmt.Println("Failed to create group : ", result.Error)
	}
	return group
}

func ListAllGroups() []models.Group {
	var groups []models.Group
	d := db.GetDB()
	d.Find(&groups)
	return groups
}

func UpdateGroup(g models.Group, Id string) *models.Group {
	d := db.GetDB()
	id, err := uuid.Parse(Id)
	if err != nil {
		fmt.Println("Failed to parse uuid")
	}
	g.ID = id
	result := d.Save(&g)
	if result.Error != nil {
		fmt.Println("Failed to update group : ", result.Error)
	}
	return &g
}

func RemoveGroup(Id string) *Response {
	d := db.GetDB()
	id, err := uuid.Parse(Id)
	if err != nil {
		fmt.Println("Failed to parse uuid")
	}
	group := &models.Group{
		ID: id,
	}
	result := d.Delete(&group)
	fmt.Println(result.Logger)
	if result.Error != nil {
		fmt.Println("Failed to delete group : ", result.Error)
	}
	return &Response{
		Message: "Group deleted successfully.",
	}
}
