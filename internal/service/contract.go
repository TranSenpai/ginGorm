package service

import (
	"context"
	"encoding/base64"
	"errors"
	"main/internal/entity"
	model "main/internal/models"
	repo "main/internal/repo"
	errorx "main/internal/utils/myerror"
	"math/rand/v2"
	"net/http"
	"time"
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

func DecodeAvatar(avatar string) (*string, error) {
	decodedAvatar, err := base64.StdEncoding.DecodeString(avatar)
	if err != nil {
		return nil, errorx.NewMyError(http.StatusInternalServerError, "Can not parse avatar", err, time.Now())
	}
	avatarString := string(decodedAvatar)

	return &avatarString, nil
}

func (c *contractService) CreateContract(ctx context.Context, contract *model.Contract) error {
	var contractEntity *entity.Contract
	if contract == nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Invalid data format", errors.New("contract empty"), time.Now())
	}

	if err := CheckRequiredField(contract); err != nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Invalid data format", errors.New("contract empty"), time.Now())
	}
	contractEntity = ToEntity(contract)

	if contract.Avatar != nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Invalid Avatar format", errors.New("avatar's format is not base64"), time.Now())
	}

	avatarString, err := DecodeAvatar(*contract.Avatar)
	if err != nil {
		return err
	}
	contractEntity.Avatar = avatarString

	timeNow := time.Now()
	contractEntity.ID = uint(rand.UintN(uint(timeNow.Nanosecond())))

	return c.contractRepo.CreateContract(ctx, contractEntity)
}

func (c *contractService) UpdateContract(ctx context.Context, filter model.Filter, contract *model.Contract) error {
	var contractEntity entity.Contract
	if contract == nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Invalid Avatar format", errors.New("contract empty"), time.Now())
	}

	if contract.Avatar != nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Invalid Avatar format", errors.New("avatar's format is not base64"), time.Now())
	}
	avatarString, err := DecodeAvatar(*contract.Avatar)
	if err != nil {
		return err
	}
	contractEntity.Avatar = avatarString

	return c.contractRepo.UpdateContract(ctx, filter, &contractEntity)
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
