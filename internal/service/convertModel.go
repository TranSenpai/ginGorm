package service

import (
	"main/internal/entity"

	model "main/internal/models"
)

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
	contractModel.ID = &contractEntity.ID
	contractModel.FirstName = &contractEntity.FirstName
	contractModel.LastName = &contractEntity.LastName
	contractModel.StudentCode = &contractEntity.StudentCode
	contractModel.Email = &contractEntity.Email
	contractModel.Sign = &contractEntity.Sign
	contractModel.Phone = &contractEntity.Phone
	contractModel.IsActive = &contractEntity.IsActive

	return &contractModel
}

func (c *contractService) MapField(contract *model.Contract) (map[string]any, error) {
	mapField := make(map[string]any)

	if contract.MiddleName != nil {
		mapField["MiddleName"] = contract.MiddleName
	}
	if contract.Gender != nil {
		mapField["Gender"] = contract.Gender
	}
	if contract.Avatar != nil {
		avatarString, err := DecodeBase64(*contract.Avatar)
		if err != nil {
			return nil, err
		}
		mapField["Avatar"] = avatarString
	}
	if contract.Address != nil {
		mapField["Address"] = contract.Address
	}
	if contract.DOB != nil {
		mapField["DOB"] = contract.DOB
	}
	if contract.RoomID != nil {
		mapField["RoomID"] = contract.RoomID
	}
	if contract.NotificationChannels != nil {
		mapField["NotificationChannels"] = contract.NotificationChannels
	}
	if contract.ID != nil {
		mapField["ID"] = contract.ID
	}

	if contract.FirstName != nil {
		mapField["FirstName"] = contract.FirstName
	}

	if contract.LastName != nil {
		mapField["LastName"] = contract.LastName
	}

	if contract.StudentCode != nil {
		mapField["StudentCode"] = contract.StudentCode
	}

	if contract.Email != nil {
		mapField["Email"] = contract.Email
	}

	if contract.Sign != nil {
		mapField["Sign"] = contract.Sign
	}

	if contract.Phone != nil {
		mapField["Phone"] = contract.Phone
	}

	if contract.IsActive != nil {
		mapField["IsActive"] = contract.IsActive
	}

	return mapField, nil
}
