package repo

import (
	"main/internal/entity"
	errorx "main/internal/utils/myerror"

	"gorm.io/gorm"
)

type deleteContractTx struct {
	StudentCode any
}

func (d deleteContractTx) Execute(tx *gorm.DB) error {
	return errorx.WrapError(
		tx.Model(&entity.Contract{}).Delete(d.StudentCode).Error,
		errorx.StatusInternalServerError,
		"Server error while deleting contract")
}
