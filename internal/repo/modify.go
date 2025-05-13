package repo

import (
	"context"
	"errors"
	"main/internal/entity"
	"main/internal/models"
	errorx "main/internal/utils/myerror"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type contractRepo struct {
	db *gorm.DB
}

func (cr *contractRepo) CreateContract(ctx context.Context, createContract *entity.Contract) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Model(&entity.Contract{}).Create(createContract).WithContext(ctx).Error
		if err != nil {
			return GetError(err)
		}
		return nil
	})
}

func buildWhere(filter models.Filter, tx *gorm.DB) *gorm.DB {
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
		tx = tx.Where("gender IN ?", filter.Gender)
	}
	if filter.IsActive != nil {
		tx = tx.Where("is_active = ?", *filter.IsActive)
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

	return tx
}

func (cr *contractRepo) UpdateContract(ctx context.Context, filter models.Filter, contract *entity.Contract) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		tx = buildWhere(filter, tx)
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
		tx = buildWhere(filter, tx)
		err := tx.Debug().Model(&entity.Contract{}).Delete(&entity.Contract{}).WithContext(ctx).Error
		if err != nil {
			return GetError(err)
		}
		return nil
	})
}

func (cr *contractRepo) Search(ctx context.Context, filter models.Filter) ([]entity.Contract, error) {
	var lst []entity.Contract
	err := cr.db.Transaction(func(tx *gorm.DB) error {
		tx = buildWhere(filter, tx)
		return tx.Debug().Model(&entity.Contract{}).Find(&lst).Error
	})
	if len(lst) == 0 {
		return lst, errorx.NewMyError(http.StatusNotFound, "Your contract does not exist", errors.New("contract not found"), time.Now())
	}
	if err != nil {
		return lst, errorx.NewMyError(http.StatusInternalServerError, "Server error while searching contract", err, time.Now())
	}
	return lst, nil
}
