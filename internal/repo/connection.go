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

	dbConn.Model(&model.Contract{}).Exec(`PARTITION BY LIST COLUMNS(ID, registry_month) (
											PARTITION p_month_1 VALUES IN (1),
											PARTITION p_month_2 VALUES IN (2),
											PARTITION p_month_3 VALUES IN (3),
											PARTITION p_month_4 VALUES IN (4),
											PARTITION p_month_5 VALUES IN (5),
											PARTITION p_month_6 VALUES IN (6),
											PARTITION p_month_7 VALUES IN (7),
											PARTITION p_month_8 VALUES IN (8),
											PARTITION p_month_9 VALUES IN (9),
											PARTITION p_month_10 VALUES IN (10),
											PARTITION p_month_11 VALUES IN (11),
											PARTITION p_month_12 VALUES IN (12))`)
}

// Singleton pattern
func GetConnection() (*gorm.DB, error) {
	dsn := "root:SyChuong241203$@tcp(127.0.0.1:3306)/gorm?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
