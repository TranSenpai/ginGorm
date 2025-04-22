package repo

import (
	model "main/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type struct DB
var (
	dbConn *gorm.DB
)

func init() {
	data, err := GetConnection()
	if err != nil {
		panic(err)
	}
	dbConn = data

	dbConn.AutoMigrate(&model.Contract{})
	dbConn.AutoMigrate(&model.Role{})
	dbConn.AutoMigrate(&model.Room{})
}

// Singleton pattern
func GetConnection() (*gorm.DB, error) {
	dsn := "root:SyChuong241203$@tcp(127.0.0.1:3306)/gorm?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
