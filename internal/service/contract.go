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

type IContractService interface {
	CreateContract(ctx context.Context, contract *model.Contract) error
	UpdateContract(ctx context.Context, filter model.Filter, contract *model.Contract) error
	DeleteContract(ctx context.Context, filter model.Filter) error
	Search(ctx context.Context, filter model.Filter) ([]model.Contract, error)
}

type contractService struct {
	contractRepo repo.IContractRepo
}

var ContractService *contractService

func NewContractService() *contractService {
	if ContractService == nil {
		ContractService = &contractService{contractRepo: repo.GetInstanceContract()}
	}
	return ContractService
}

func (c *contractService) GetContractService() IContractService {
	return ContractService
}

func ToEntity(contractModel *model.Contract) *entity.Contract {
	var contractEntity entity.Contract
	if contractModel.FirstName != nil {
		contractEntity.FirstName = contractModel.FirstName
	}
	if contractModel.LastName != nil {
		contractEntity.LastName = contractModel.LastName
	}
	if contractModel.MiddleName != nil {
		contractEntity.MiddleName = contractModel.MiddleName
	}
	if contractModel.Gender != nil {
		contractEntity.Gender = contractModel.Gender
	}
	if contractModel.Avatar != nil {
		contractEntity.Avatar = contractModel.Avatar
	}
	if contractModel.Address != nil {
		contractEntity.Address = contractModel.Address
	}
	if contractModel.DOB != nil {
		contractEntity.DOB = contractModel.DOB
	}
	if contractModel.RoomID != nil {
		contractEntity.RoomID = contractModel.RoomID
	}
	if contractModel.IsActive != nil {
		contractEntity.IsActive = contractModel.IsActive
	}
	if contractModel.NotificationChannels != nil {
		contractEntity.NotificationChannels = contractModel.NotificationChannels
	}

	contractEntity.StudentCode = contractModel.StudentCode
	contractEntity.Email = contractModel.Email
	contractEntity.Sign = contractModel.Sign
	contractEntity.Phone = contractModel.Phone

	return &contractEntity
}

func ToContract(contractEntity *entity.Contract) *model.Contract {
	var contractModel model.Contract
	if contractEntity.FirstName != nil {
		contractModel.FirstName = contractEntity.FirstName
	}
	if contractEntity.LastName != nil {
		contractModel.LastName = contractEntity.LastName
	}
	if contractEntity.MiddleName != nil {
		contractModel.MiddleName = contractEntity.MiddleName
	}
	if contractEntity.Gender != nil {
		contractModel.Gender = contractEntity.Gender
	}
	if contractEntity.Avatar != nil {
		contractModel.Avatar = contractEntity.Avatar
	}
	if contractEntity.Address != nil {
		contractModel.Address = contractEntity.Address
	}
	if contractEntity.DOB != nil {
		contractModel.DOB = contractEntity.DOB
	}
	if contractEntity.RoomID != nil {
		contractModel.RoomID = contractEntity.RoomID
	}
	if contractEntity.IsActive != nil {
		contractModel.IsActive = contractEntity.IsActive
	}
	if contractEntity.NotificationChannels != nil {
		contractModel.NotificationChannels = contractEntity.NotificationChannels
	}

	contractModel.StudentCode = contractEntity.StudentCode
	contractModel.Email = contractEntity.Email
	contractModel.Sign = contractEntity.Sign
	contractModel.Phone = contractEntity.Phone

	return &contractModel
}

func (c *contractService) CreateContract(ctx context.Context, contract *model.Contract) error {
	contractEntity := ToEntity(contract)
	if contract.Avatar != nil {
		decodedAvatar, err := base64.StdEncoding.DecodeString(*contract.Avatar)
		if err != nil {
			return errorx.New(http.StatusUnprocessableEntity, "Invalid Avatar format", err)
		}
		avatarString := string(decodedAvatar)
		contractEntity.Avatar = &avatarString
	}

	return c.contractRepo.CreateContract(ctx, contractEntity)
}

func (c *contractService) UpdateContract(ctx context.Context, filter model.Filter, contract *model.Contract) error {
	contractEntity := ToEntity(contract)
	if contract.Avatar != nil {
		decodedAvatar, err := base64.StdEncoding.DecodeString(*contract.Avatar)
		if err != nil {
			return errorx.New(http.StatusUnprocessableEntity, "Invalid Avatar format", err)
		}
		avatarString := string(decodedAvatar)
		contractEntity.Avatar = &avatarString
	}

	return c.contractRepo.UpdateContract(ctx, filter, contractEntity)
}

func (c *contractService) DeleteContract(ctx context.Context, filter model.Filter) error {
	return c.contractRepo.DeleteContract(ctx, filter)
}

func (s contractService) Search(ctx context.Context, filter model.Filter) ([]model.Contract, error) {
	entities, err := s.contractRepo.Search(ctx, filter)

	if err != nil {
		return nil, err
	}

	var contracts []model.Contract
	for _, v := range entities {
		contracts = append(contracts, *ToContract(&v))
	}

	return contracts, nil
}
