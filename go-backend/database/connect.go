package database

import (
	"go-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//Connect to database
	database, err := gorm.Open(mysql.Open("root:sc4d4@/mydb?parseTime=true"), &gorm.Config{})

	//check database connection
	if err != nil {
		panic("Could not connect to database")
	}

	DB = database
	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}
