package service

import (
	model "main/internal/models"
	repo "main/internal/repo"
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

func (c contractService) CreateContract(contract *model.Contract) error {
	contractRepo := repo.GetInstanceContract()

	err := contractRepo.RegisterContract(contract)
	if err != nil {
		return err
	}

	return nil
}

func (c contractService) DeleteContract(id string) error {
	contractRepo := repo.GetInstanceContract()

	err := contractRepo.DeleteContract(id)
	if err != nil {
		return err
	}

	return nil
}

func (c contractService) UpdateContract(id string, contract *model.Contract) error {
	contractRepo := repo.GetInstanceContract()

	err := contractRepo.UpdateContract(id, contract)
	if err != nil {
		return err
	}

	return nil
}

func (c contractService) Search(id string) (model.Contract, error) {
	contractRepo := repo.GetInstanceContract()

	contract, err := contractRepo.Search(id)
	if err != nil {
		return model.Contract{}, err
	}
	return contract, nil
}

func (c contractService) SearchAll() ([]model.Contract, error) {
	contractRepo := repo.GetInstanceContract()

	contract, err := contractRepo.SearchAll()
	if err != nil {
		return nil, err
	}
	return contract, nil
}
