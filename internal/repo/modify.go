package repo

import (
	"context"
	"errors"
	"main/internal/entity"
	errorx "main/internal/utils/myerror"

	"gorm.io/gorm"
)

type contractRepo struct {
	db *gorm.DB
}

func (cr contractRepo) CreateContract(ctx context.Context, createContract *entity.Contract) error {
	err := CheckIDTransaction(ctx)
	if err != nil {
		return err
	}
	return RunTransaction(ctx, cr.db, createContractTx{Contract: createContract})
}

func (cr contractRepo) UpdateContract(ctx context.Context, studentCode string, updateContract *entity.Contract) error {
	err := CheckIDTransaction(ctx)
	if err != nil {
		return err
	}
	return RunTransaction(ctx, cr.db, updateContractTx{Contract: updateContract, StudentCode: studentCode})
}

func (cr contractRepo) DeleteContract(ctx context.Context, studentCode string) error {
	err := CheckIDTransaction(ctx)
	if err != nil {
		return err
	}
	return RunTransaction(ctx, cr.db, deleteContractTx{StudentCode: studentCode})
}

func (cr contractRepo) Search(ctx context.Context, studentCode string) (entity.Contract, error) {
	var contract entity.Contract
	err := cr.db.WithContext(ctx).Where("student_code = ?", studentCode).First(&contract).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return contract, errorx.WrapError(err, errorx.StatusNotFound, "Contract not found")
		}
		return contract, errorx.WrapError(err, errorx.StatusInternalServerError, "Server error while searching contract")
	}
	return contract, nil
}

func (cr contractRepo) SearchAll() ([]entity.Contract, error) {
	var contracts []entity.Contract
	err := cr.db.Find(&contracts).Error
	return contracts, errorx.WrapError(err, errorx.StatusInternalServerError, "Server error while searching all contracts")
}
