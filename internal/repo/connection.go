package repo

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type struct DB
var (
	dbConn *gorm.DB
)

func getConnection() {
	data, err := buildConnection()
	if err != nil {
		panic(err)
	}
	dbConn = data
}

func buildConnection() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Cannot load .env file, fallback to OS environment")
	}

	username := os.Getenv("ROOT_USERNAME")
	password := os.Getenv("ROOT_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	if username == "" || password == "" || host == "" || port == "" || dbname == "" {
		return nil, fmt.Errorf("missing database configuration")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbname)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{QueryFields: true})
}
