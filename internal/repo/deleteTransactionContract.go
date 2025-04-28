package repo

import (
	"main/internal/entity"
	errorx "main/internal/utils/myerror"

	"gorm.io/gorm"
)

type deleteContractTx struct {
	StudentCode any
}

func (d deleteContractTx) Execute(db *gorm.DB) error {
	err := db.Model(&entity.Contract{}).Delete(d.StudentCode).Error
	if err != nil {
		return errorx.New(errorx.StatusInternalServerError, "Server error while deleting contract", err)
	}
	return nil
}
