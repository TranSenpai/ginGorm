package service

import (
	"main/internal/entity"
	repo "main/internal/repo"
)

type Builder interface {
	MapStudentInfo() *ContractBuilder
	MapContactInfo() *ContractBuilder
	MapRoomInfo() *ContractBuilder
	MapAvatar() *ContractBuilder
	ValidateUniqueness(repo repo.IContractRepo) *ContractBuilder
	CheckCreatedContract(repo repo.IContractRepo) *ContractBuilder
	GetContract() (*entity.Contract, error)
}
