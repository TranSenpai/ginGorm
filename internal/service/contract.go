package service

import (
	"context"
	"encoding/base64"
	"main/internal/entity"
	model "main/internal/models"
	repo "main/internal/repo"
	errorx "main/internal/utils/myerror"
	"net/http"
)

type IService interface {
	CreateContract(ctx context.Context, m *model.Contract) error
	UpdateContract(ctx context.Context, filter *model.Filter, m *model.Contract) error
	DeleteContract(ctx context.Context, filter *model.Filter) error
	Search(ctx context.Context, filter *model.Filter) ([]model.Contract, error)
}

type contractService struct {
	contractRepo repo.IRepo
}

var ContractService *contractService

func NewContractService() *contractService {
	if ContractService == nil {
		ContractService = &contractService{contractRepo: repo.GetInstanceContract()}
	}
	return ContractService
}

func toEntity(c model.Contract) *entity.Contract {
	return &entity.Contract{
		StudentCode:          c.StudentCode,
		FullName:             &c.FullName,
		Email:                c.Email,
		Sign:                 c.Sign,
		Phone:                c.Phone,
		Gender:               &c.Gender,
		DOB:                  &c.DOB,
		Address:              &c.Address,
		RoomID:               &c.RoomID,
		IsActive:             &c.IsActive,
		NotificationChannels: &c.NotificationChannels,
	}
}

func toContract(e entity.Contract) *model.Contract {
	return &model.Contract{
		StudentCode:          e.StudentCode,
		FullName:             *e.FullName,
		Email:                e.Email,
		Sign:                 e.Sign,
		Phone:                e.Phone,
		Gender:               *e.Gender,
		DOB:                  *e.DOB,
		Address:              *e.Address,
		IsActive:             *e.IsActive,
		RoomID:               *e.RoomID,
		NotificationChannels: *e.NotificationChannels,
	}
}

func (c *contractService) GetContractService() IService {
	return ContractService
}

func (c *contractService) CreateContract(ctx context.Context, m *model.Contract) error {
	decodedAvatar, err := base64.StdEncoding.DecodeString(m.Avatar)
	if err != nil {
		return errorx.New(http.StatusUnprocessableEntity, "Invalid Avatar format", err)
	}
	contract := toEntity(*m)
	avatarString := string(decodedAvatar)
	contract.Avatar = &avatarString

	return c.contractRepo.CreateContract(ctx, contract)
}

func (c *contractService) UpdateContract(ctx context.Context, filter *model.Filter, m *model.Contract) error {
	decodedAvatar, err := base64.StdEncoding.DecodeString(m.Avatar)
	if err != nil {
		return errorx.New(http.StatusUnprocessableEntity, "Invalid Avatar format", err)
	}
	contract := toEntity(*m)
	avatarString := string(decodedAvatar)
	contract.Avatar = &avatarString

	return c.contractRepo.UpdateContract(ctx, filter, contract)
}

func (c *contractService) DeleteContract(ctx context.Context, filter *model.Filter) error {
	return c.contractRepo.DeleteContract(ctx, filter)
}

func (s contractService) Search(ctx context.Context, filter *model.Filter) ([]model.Contract, error) {
	entities, err := s.contractRepo.Search(ctx, filter)

	if err != nil {
		return nil, err
	}

	var contracts []model.Contract
	for _, v := range entities {
		contracts = append(contracts, *toContract(v))
	}

	return contracts, nil
}
