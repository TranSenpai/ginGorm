package repo

import "gorm.io/gorm"

// type struct DB
var dbConnection *gorm.DB

// var dbConnection2 *gorm.DB

func init() {
	mysqlConnection := NewMySQL()
	var err error
	dbConnection, err = mysqlConnection.GetConnection(dbConnection)
	if err != nil {
		panic(err)
	}
	createTableContract(dbConnection)
	createPartitionContract(dbConnection)
}
