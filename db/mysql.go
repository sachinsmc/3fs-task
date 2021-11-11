package db

import (
	"fmt"
	"github.com/sachinsmc/3fs-task/models"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	host := os.Getenv("DATABASE_HOST")
	if host != "" {
		host = os.Getenv("db.host")
	}
	dbname := viper.GetString("db.name")
	user := viper.GetString("db.username")
	password := viper.GetString("db.password")
	port := viper.GetString("db.port")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database!")
		panic(err)
	} else {
		fmt.Println("Connected to database.")
	}

	DB.Debug()

	DB.Debug().AutoMigrate(&models.Group{}, &models.Users{})

}

func GetDB() *gorm.DB {
	return DB
}
