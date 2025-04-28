package repo

import (
	"context"
	"errors"
	"main/internal/entity"
	errorx "main/internal/utils/myerror"

	"gorm.io/gorm"
)

type modifyContractByName struct {
	contractRepo
}

func (m modifyContractByName) CreateContract(ctx context.Context, createContract *entity.Contract) error {
	if err := CheckIDTransaction(ctx); err != nil {
		return err
	}
	return RunTransaction(ctx, m.db, createContractTx{Contract: createContract})
}

func (m modifyContractByName) UpdateContract(ctx context.Context, studentCode string, updateContract *entity.Contract) error {
	if err := CheckIDTransaction(ctx); err != nil {
		return err
	}
	return RunTransaction(ctx, m.db, updateContractTx{Contract: updateContract, StudentCode: studentCode})
}

func (m modifyContractByName) DeleteContract(ctx context.Context, fullName string) error {
	if err := CheckIDTransaction(ctx); err != nil {
		return err
	}
	return RunTransaction(ctx, m.db, deleteContractTx{StudentCode: fullName})
}

func (m modifyContractByName) Search(ctx context.Context, fullName string) (entity.Contract, error) {
	var contract entity.Contract
	err := m.db.WithContext(ctx).Where("full_name = ?", fullName).First(&contract).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return contract, errorx.New(errorx.StatusNotFound, "Contract not found by full name", err)
		}
		return contract, errorx.New(errorx.StatusInternalServerError, "Server error while searching by full name", err)
	}
	return contract, nil
}

func (m modifyContractByName) SearchAll() ([]entity.Contract, error) {
	return m.contractRepo.SearchAll()
}
