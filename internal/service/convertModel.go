package service

import (
	"main/internal/entity"
	model "main/internal/models"
)

func ToEntity(contractModel *model.Contract) *entity.Contract {
	var contractEntity entity.Contract
	if contractModel.StudentCode != nil {
		contractEntity.StudentCode = *contractModel.StudentCode
	}
	if contractModel.FirstName != nil {
		contractEntity.FirstName = *contractModel.FirstName
	}
	if contractModel.LastName != nil {
		contractEntity.LastName = *contractModel.LastName
	}
	if contractModel.Email != nil {
		contractEntity.Email = *contractModel.Email
	}
	if contractModel.Sign != nil {
		contractEntity.Sign = *contractModel.Sign
	}
	if contractModel.Phone != nil {
		contractEntity.Phone = *contractModel.Phone
	}
	if contractModel.IsActive != nil {
		contractEntity.IsActive = contractModel.IsActive
	}
	if contractModel.MiddleName != nil {
		contractEntity.MiddleName = contractModel.MiddleName
	}
	if contractModel.Gender != nil {
		contractEntity.Gender = contractModel.Gender
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
	if contractModel.NotificationChannels != nil {
		contractEntity.NotificationChannels = contractModel.NotificationChannels
	}

	return &contractEntity
}

func ToContract(contractEntity *entity.Contract) *model.Contract {
	var contractModel model.Contract
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
	if contractEntity.NotificationChannels != nil {
		contractModel.NotificationChannels = contractEntity.NotificationChannels
	}
	*contractModel.ID = contractEntity.ID
	*contractModel.FirstName = contractEntity.FirstName
	*contractModel.LastName = contractEntity.LastName
	*contractModel.StudentCode = contractEntity.StudentCode
	*contractModel.Email = contractEntity.Email
	*contractModel.Sign = contractEntity.Sign
	*contractModel.Phone = contractEntity.Phone
	*contractModel.IsActive = *contractEntity.IsActive

	return &contractModel
}
