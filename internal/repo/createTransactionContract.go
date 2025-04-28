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
	return errorx.WrapError(tx.Debug().Create(&c.Contract).Error,
		errorx.StatusInternalServerError,
		"Server error while creating contract")
}
