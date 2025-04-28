package repo

import (
	"main/internal/entity"
	errorx "main/internal/utils/myerror"

	"gorm.io/gorm"
)

type updateContractTx struct {
	Contract    *entity.Contract
	StudentCode string
}

func (u updateContractTx) Execute(db *gorm.DB) error {
	return errorx.WrapError(
		db.Debug().Model(&entity.Contract{}).
			Where("student_code = ?", u.StudentCode).
			Updates(&u.Contract).Error,
		errorx.StatusInternalServerError,
		"Server error while updating contract",
	)
}
