package repo

import (
	"main/internal/entity"
	errorx "main/internal/utils/myerror"

	"gorm.io/gorm"
)

type updateContractTx struct {
	Contract *entity.Contract
}

func (u updateContractTx) Execute(db *gorm.DB) error {
	err := db.Model(&entity.Contract{}).Updates(&u.Contract).Error
	if err != nil {
		return errorx.New(errorx.StatusInternalServerError, "Server error while updating contract", err)
	}
	return nil
}
