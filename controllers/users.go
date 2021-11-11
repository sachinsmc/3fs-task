package controllers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sachinsmc/3fs-task/db"
	"github.com/sachinsmc/3fs-task/models"
	"golang.org/x/crypto/bcrypt"
)

func InsertUser(u *models.Users) *models.Users {
	d := db.GetDB()
	u.ID = uuid.New()
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hash)
	result := d.Create(u)
	if result.Error != nil {
		fmt.Println("Failed to create user : ", result.Error)
	}
	u.Password = ""
	return u
}

func ListAllUsers() []models.Users {
	var users []models.Users
	d := db.GetDB()
	d.Find(&users)
	for i, _ := range users {
		users[i].Password = ""
	}
	return users
}

func UpdateUser(u models.Users, Id string) *models.Users {
	d := db.GetDB()
	id, err := uuid.Parse(Id)
	if err != nil {
		fmt.Println("Failed to parse uuid")
	}
	u.ID = id
	d.Model(models.Users{}).Where("id = ?", id).Update("name", u.Name)
	return &u
}

func RemoveUser(Id string) *Response {
	d := db.GetDB()
	id, err := uuid.Parse(Id)
	if err != nil {
		fmt.Println("Failed to parse uuid")
	}
	user := &models.Users{
		ID: id,
	}
	result := d.Delete(&user)
	if result.Error != nil {
		fmt.Println("Failed to delete user : ", result.Error)
	}
	return &Response{
		Message: "User deleted successfully.",
	}
}
