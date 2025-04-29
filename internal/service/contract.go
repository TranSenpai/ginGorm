package service

import (
	"context"
	"errors"
	"main/internal/entity"
	model "main/internal/models"
	repo "main/internal/repo"
	errorx "main/internal/utils/myerror"
)

type contractService struct{}

var (
	ContractService *contractService
)

func GetContractService() *contractService {
	if ContractService == nil {
		ContractService = &contractService{}
	}
	return ContractService
}

func (c contractService) CreateContract(ctx context.Context, m *model.Contract) error {
	contractRepo := repo.GetInstanceContract()
	buildContract := NewContractBuilder(ctx, m)
	buildContract = buildContract.MapStudentInfo().MapContactInfo().MapRoomInfo().MapAvatar()
	if buildContract.buildError != nil {
		return buildContract.buildError
	}

	if _, err := c.Search(ctx, m.StudentCode); err == nil {
		return errorx.WrapError(errors.New("Contract has benn existed"), errorx.StatusConflict, "Dupplicate contract")
	}

	ctx = context.WithValue(ctx, "id", buildContract.entity.StudentCode)

	return contractRepo.CreateContract(ctx, buildContract.entity)
}

func (c contractService) UpdateContract(ctx context.Context, studentCode string, m *model.Contract) error {
	contractRepo := repo.GetInstanceContract()
	if _, err := contractRepo.Search(ctx, studentCode); err != nil {
		return err
	}

	buildContract := NewContractBuilder(ctx, m)
	buildContract = buildContract.MapStudentInfo().MapContactInfo().MapRoomInfo().MapAvatar()
	if buildContract.buildError != nil {
		return buildContract.buildError
	}
	ctx = context.WithValue(ctx, "id", studentCode)

	return contractRepo.UpdateContract(ctx, studentCode, buildContract.entity)
}

func (c contractService) DeleteContract(ctx context.Context, studentCode string) error {
	contractRepo := repo.GetInstanceContract()
	if _, err := contractRepo.Search(ctx, studentCode); err != nil {
		return err
	}

	ctx = context.WithValue(ctx, "id", studentCode)
	return contractRepo.DeleteContract(ctx, studentCode)
}

func (c contractService) Search(ctx context.Context, studentCode string) (entity.Contract, error) {
	contractRepo := repo.GetInstanceContract()
	return contractRepo.Search(ctx, studentCode)
}

func (c contractService) SearchAll() ([]entity.Contract, error) {
	contractRepo := repo.GetInstanceContract()
	return contractRepo.SearchAll()
}

func (c contractService) SearchByName(ctx context.Context, fullName string) (entity.Contract, error) {
	contractRepo := repo.GetModifyContractByName()
	return contractRepo.Search(ctx, fullName)
}
