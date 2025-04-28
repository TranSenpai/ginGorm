package service

import (
	"context"
	"encoding/base64"
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

func toEntityContract(m *model.Contract) (*entity.Contract, error) {
	avatarData, err := base64.StdEncoding.DecodeString(m.Avatar)
	if err != nil {
		return nil, errorx.New(errorx.StatusUnprocessableEntity, "Invalid Avatar format", err)
	}

	entityContract := &entity.Contract{
		StudentCode:          m.StudentCode,
		FullName:             m.FullName,
		Email:                m.Email,
		Sign:                 m.Sign,
		Phone:                m.Phone,
		Gender:               m.Gender,
		DOB:                  &m.DOB,
		Address:              &m.Address,
		IsActive:             m.IsActive,
		RoomID:               m.RoomID,
		NotificationChannels: m.NotificationChannels,
		Avatar:               avatarData,
	}

	return entityContract, nil
}

func (c contractService) CreateContract(ctx context.Context, m *model.Contract) error {
	contractRepo := repo.GetInstanceContract()

	builder := NewContractBuilder(ctx, m)
	// MapStudentInfo().
	// MapContactInfo().
	// MapRoomInfo().
	// MapAvatar()
	builder.MapStudentInfo()
	if builder.buildError != nil {
		return builder.buildError
	}

	builder.ValidateUniqueness(contractRepo)
	if builder.buildError != nil {
		return builder.buildError
	}

	entityContract, err := builder.GetContract()
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, "id", entityContract.StudentCode)

	return contractRepo.CreateContract(ctx, entityContract)
}

func (c contractService) UpdateContract(ctx context.Context, studentCode string, m *model.Contract) error {
	contractRepo := repo.GetInstanceContract()

	_, err := contractRepo.Search(ctx, studentCode)
	if err != nil {
		if myerr, ok := err.(errorx.MyErr); ok {
			return errorx.New(myerr.Status, "Cannot update non-existing contract", myerr.Err)
		}
		return errorx.New(errorx.StatusInternalServerError, "Failed to search contract", err)
	}

	entityContract, err := toEntityContract(m)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, "id", studentCode)
	return contractRepo.UpdateContract(ctx, entityContract)
}

func (c contractService) DeleteContract(ctx context.Context, studentCode string) error {
	contractRepo := repo.GetInstanceContract()

	_, err := contractRepo.Search(ctx, studentCode)
	if err != nil {
		if myerr, ok := err.(errorx.MyErr); ok {
			return errorx.New(myerr.Status, "Cannot delete non-existing contract", myerr.Err)
		}
		return errorx.New(errorx.StatusInternalServerError, "Failed to search contract", err)
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
