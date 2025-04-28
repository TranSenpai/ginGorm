package entity

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type struct DB
var (
	DBConn *gorm.DB
	once   sync.Once // Khai ở global để bảm bảo chạy 1 lần trong 1 vòng đời program

)

// Singleton pattern
func GetConnection() (*gorm.DB, error) {
	dsn := "root:SyChuong241203$@tcp(127.0.0.1:3306)/gorm?parseTime=true"
	var err error
	if DBConn == nil {
		once.Do(
			func() {
				DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
				if err != nil {
					log.Fatalln("Failed to connect DB:", err)
				}
			})
	} else {
		fmt.Println("Database connection is created")
	}
	return DBConn, err
}
