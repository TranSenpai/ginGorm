package repo

import (
	"context"
	"errors"
	"main/internal/entity"
	"main/internal/models"

	"gorm.io/gorm"
)

type IContractRepo interface {
	CreateContract(ctx context.Context, contract *entity.Contract) error
	UpdateContract(ctx context.Context, filter models.Filter, contract *entity.Contract) error
	DeleteContract(ctx context.Context, filter models.Filter) error
	Search(ctx context.Context, filter models.Filter) ([]entity.Contract, error)
	GetTotalContractEachRoom(ctx context.Context) ([]models.TotalContracts, error)
}

var ContractRepo *contractRepo

func GetInstanceContract() IContractRepo {
	if ContractRepo == nil {
		ContractRepo = &contractRepo{db: dbConnection}
	}
	if ContractRepo.db == nil {
		panic(errors.New("errorrrr"))
	}
	return ContractRepo
}

type contractRepo struct {
	db *gorm.DB
}

func (cr *contractRepo) CreateContract(ctx context.Context, createContract *entity.Contract) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Create(createContract).WithContext(ctx).Error
		if err != nil {
			return GetError(err)
		}
		return nil
	})
}

func (cr *contractRepo) buildWhere(filter models.Filter, tx *gorm.DB) *gorm.DB {
	if filter.ID != nil {
		tx = tx.Where("id = ?", filter.ID)
	}
	if filter.StudentCode != nil {
		tx = tx.Where("student_code IN ?", filter.StudentCode)
	}
	if filter.Email != nil {
		tx = tx.Where("email IN ?", filter.Email)
	}
	if filter.FirstName != nil {
		tx = tx.Where("first_name like ?", *filter.FirstName+"%")
	}
	if filter.LastName != nil {
		tx = tx.Where("last_name like ?", *filter.LastName+"%")
	}
	if filter.MiddleName != nil {
		tx = tx.Where("middle_name like ?", *filter.MiddleName+"%")
	}
	if filter.Phone != nil {
		tx = tx.Where("phone IN ?", filter.Phone)
	}
	if filter.Sign != nil {
		tx = tx.Where("sign IN ?", filter.Sign)
	}
	if filter.RoomID != nil {
		tx = tx.Where("room_id IN ?", filter.RoomID)
	}
	if filter.DOB != nil {
		tx = tx.Where("dob IN ?", filter.DOB)
	}
	if filter.Gender != nil {
		tx = tx.Where("gender = ?", filter.Gender)
	}
	if filter.NotificationChannels != nil {
		tx = tx.Where("notification_channels IN ?", filter.NotificationChannels)
	}
	if filter.Address != nil {
		tx = tx.Where("address IN ?", filter.Address)
	}
	if filter.Avatar != nil {
		tx = tx.Where("avatar IN ?", filter.Avatar)
	}
	if filter.IsActive != nil {
		tx = tx.Where("is_active = ?", filter.IsActive)
	}

	return tx
}

func (cr *contractRepo) UpdateContract(ctx context.Context, filter models.Filter, contract *entity.Contract) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		tx = cr.buildWhere(filter, tx)
		// Updates supports updating with struct or map[string]interface{},
		// when updating with struct it will only update non-zero fields by default
		err := tx.Debug().Model(&entity.Contract{}).Updates(*contract).WithContext(ctx).Error
		if err != nil {
			return GetError(err)
		}
		return nil
	})
}

func (cr *contractRepo) DeleteContract(ctx context.Context, filter models.Filter) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		tx = cr.buildWhere(filter, tx)
		err := tx.Debug().Delete(&entity.Contract{}).WithContext(ctx).Error
		if err != nil {
			return GetError(err)
		}
		return nil
	})
}

func (cr *contractRepo) Search(ctx context.Context, filter models.Filter) ([]entity.Contract, error) {
	var lst []entity.Contract
	cr.db = cr.buildWhere(filter, cr.db)
	err := cr.db.Debug().Model(&entity.Contract{}).Find(&lst).Error
	if err != nil {
		GetError(err)
	}

	return lst, err
}

func (cr *contractRepo) GetTotalContractEachRoom(ctx context.Context) ([]models.TotalContracts, error) {
	var result []models.TotalContracts
	err := cr.db.Debug().Model(&entity.Contract{}).
		Select("COUNT(id) as total, room_id").Where("is_active = ?", true).
		Group("room_id").Having("total >= ?", 2).Find(&result).Error
	if err != nil {
		GetError(err)
	}

	return result, err
}
