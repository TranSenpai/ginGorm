package repo

import (
	"fmt"
	"log"
	"main/internal/entity"
	"os"

	"github.com/joho/godotenv"
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

	err = CreateContractTable(dbConn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dbConn.Debug().Model(&entity.Contract{}).Exec(`
	ALTER TABLE gorm.contracts
	ADD COLUMN id INT UNSIGNED AUTO_INCREMENT,
  	ADD COLUMN registry_month TINYINT GENERATED ALWAYS AS (MONTH(registry_at)) STORED`)

	dbConn.Debug().Model(&entity.Contract{}).Exec(`
	ALTER TABLE gorm.contracts
	PARTITION BY LIST COLUMNS (registry_month) (
		PARTITION p01 VALUES IN (1),
		PARTITION p02 VALUES IN (2),
		PARTITION p03 VALUES IN (3),
		PARTITION p04 VALUES IN (4),
		PARTITION p05 VALUES IN (5),
		PARTITION p06 VALUES IN (6),
		PARTITION p07 VALUES IN (7),
		PARTITION p08 VALUES IN (8),
		PARTITION p09 VALUES IN (9),
		PARTITION p10 VALUES IN (10),
		PARTITION p11 VALUES IN (11),
		PARTITION p12 VALUES IN (12)
	)`)
}

func GetConnection() (*gorm.DB, error) {
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
