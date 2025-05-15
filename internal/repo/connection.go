package repo

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type IConnection interface {
	GetConnection(dbConnection *gorm.DB)
}

type MySQL struct{}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func (m *MySQL) GetConnection(dbConnection *gorm.DB) (*gorm.DB, error) {
	if dbConnection == nil {
		return buildMySQLConnection()
	}
	return dbConnection, nil
}

func buildMySQLConnection() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Cannot load .env file, fallback to OS environment")
	}

	username := os.Getenv("MySQL_ROOT_USERNAME")
	password := os.Getenv("MySQL_ROOT_PASSWORD")
	host := os.Getenv("MySQL_DB_HOST")
	port := os.Getenv("MySQL_DB_PORT")
	dbname := os.Getenv("MySQL_DB_NAME")

	if username == "" || password == "" || host == "" || port == "" || dbname == "" {
		return nil, fmt.Errorf("missing database configuration")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbname)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{QueryFields: true})
}
