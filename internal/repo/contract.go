package repo

import (
	"context"
	"errors"
	"main/internal/entity"
	errorx "main/internal/utils/myerror"

	"gorm.io/gorm"
)

var (
	ContractRepo         *contractRepo
	ModifyContractByName *modifyContractByName
)

func CheckIDTransaction(ctx context.Context) error {
	id := ctx.Value("id")
	if id == nil {
		return errorx.New(errorx.StatusUnauthorized, "Missing ID Transaction", errors.New("missing id transaction"))
	}
	return nil
}

func CreateContractTable(connection *gorm.DB) error {
	return connection.AutoMigrate(&entity.Contract{})
}

func GetInstanceContract() IContractRepo {
	if ContractRepo == nil {
		ContractRepo = &contractRepo{db: dbConn}
	}
	return ContractRepo
}

func GetModifyContractByName() IContractRepo {
	if ModifyContractByName == nil {
		ModifyContractByName = &modifyContractByName{contractRepo{db: dbConn}}
	}
	return ModifyContractByName
}
