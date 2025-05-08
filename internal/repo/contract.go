package repo

import (
	"main/internal/entity"

	"gorm.io/gorm"
)

var ContractRepo *contractRepo

func CreateContractTable(connection *gorm.DB) error {
	return connection.AutoMigrate(&entity.Contract{})
}

func GetInstanceContract() IRepo {
	if ContractRepo == nil {
		ContractRepo = &contractRepo{db: dbConn}
	}
	return ContractRepo
}
