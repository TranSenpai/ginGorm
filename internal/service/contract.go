package service

import (
	"context"
	"encoding/base64"
	"errors"
	"main/internal/entity"
	model "main/internal/models"
	"main/internal/repo"

	errorx "main/internal/utils/myerror"
	"net/http"
	"time"
)

type IContractService interface {
	CreateContract(ctx context.Context, contract *model.Contract) error
	UpdateContract(ctx context.Context, contractID uint, contract *model.Contract) error
	DeleteContract(ctx context.Context, filter model.Filter) error
	Search(ctx context.Context, filter model.Filter) ([]model.Contract, error)
	GetTotalContractEachRoom(ctx context.Context) ([]model.TotalContracts, error)
	SignContract(ctx context.Context, filter model.Filter, signature string) error
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

func DecodeBase64(input string) (*string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, errorx.NewMyError(http.StatusInternalServerError, "Can not parse avatar", err, time.Now())
	}
	result := string(decoded)

	return &result, nil
}

func (c *contractService) CreateContract(ctx context.Context, contract *model.Contract) error {
	if contract == nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Invalid data format", errors.New("contract empty"), time.Now())
	}

	if err := c.CheckRequiredField(ctx, contract); err != nil {
		return err
	}

	mapField, err := c.MapField(contract)
	if err != nil {
		return err
	}
	mapField["IsActive"] = false
	mapField["Sign"] = ""

	return c.contractRepo.CreateContract(ctx, mapField)
}

func (c *contractService) UpdateContract(ctx context.Context, contractID uint, contract *model.Contract) error {
	if contract == nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Invalid data format", errors.New("nil contract"), time.Now())
	}

	mapField, err := c.MapField(contract)
	if err != nil {
		return err
	}

	return c.contractRepo.UpdateContract(ctx, contractID, mapField)
}

func (c *contractService) DeleteContract(ctx context.Context, filter model.Filter) error {
	return c.contractRepo.DeleteContract(ctx, filter)
}

func (c *contractService) Search(ctx context.Context, filter model.Filter) ([]model.Contract, error) {
	entities, err := c.contractRepo.Search(ctx, filter)

	if err != nil {
		return nil, err
	}

	var contracts []model.Contract
	for _, v := range entities {
		contracts = append(contracts, *ToContract(&v))
	}

	return contracts, nil
}

func (c *contractService) GetTotalContractEachRoom(ctx context.Context) ([]model.TotalContracts, error) {
	totalContract, err := c.contractRepo.GetTotalContractEachRoom(ctx)
	if err != nil {
		return nil, err
	}

	return totalContract, nil
}

func (c *contractService) ValidateSignContract(ctx context.Context, contract *entity.Contract, signature string) error {
	if contract == nil {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Invalid data format", errors.New("nil contract"), time.Now())
	}

	if contract.IsActive {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Contract is active", errors.New("contract is active"), time.Now())
	}

	if contract.Sign != "" {
		return errorx.NewMyError(http.StatusUnprocessableEntity, "Contract is active", errors.New("contract is active"), time.Now())
	}

	contract.IsActive = true
	contract.Sign = signature

	return nil
}

func (c *contractService) SignContract(ctx context.Context, filter model.Filter, signature string) error {
	contractEntities, err := c.contractRepo.Search(ctx, filter)
	if err != nil {
		return err
	}
	if len(contractEntities) == 0 {
		return errorx.NewMyError(http.StatusNotFound, "Contract not found", errors.New("contract not found"), time.Now())
	}

	contractEntity := contractEntities[0]
	err = c.ValidateSignContract(ctx, &contractEntity, signature)
	if err != nil {
		return err
	}
	contract := ToContract(&contractEntity)
	mapField, err := c.MapField(contract)
	if err != nil {
		return err
	}

	err = c.contractRepo.UpdateContract(ctx, contractEntity.ID, mapField)
	if err != nil {
		return err
	}

	return nil
}
