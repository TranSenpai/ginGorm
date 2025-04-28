package repo

import (
	"main/internal/entity"
	errorx "main/internal/utils/myerror"

	"gorm.io/gorm"
)

type createContractTx struct {
	Contract *entity.Contract
}

func (c createContractTx) Execute(tx *gorm.DB) error {
	err := tx.Debug().Create(c.Contract).Error
	if err != nil {
		return errorx.New(errorx.StatusInternalServerError, "Server error while creating contract", err)
	}
	return nil
}
