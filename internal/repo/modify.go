package repo

import (
	"context"
	"errors"
	"main/internal/entity"
	"main/internal/models"
	errorx "main/internal/utils/myerror"
	"net/http"

	"gorm.io/gorm"
)

type contractRepo struct {
	db *gorm.DB
}

func (cr *contractRepo) CreateContract(ctx context.Context, createContract *entity.Contract) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		return tx.Debug().Model(&entity.Contract{}).Create(&createContract).WithContext(ctx).Error
	})
}

func (cr *contractRepo) UpdateContract(ctx context.Context, filter *models.Filter, updateContract *entity.Contract) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		var lst []entity.Contract
		lst, err := cr.Search(ctx, filter)
		if err != nil {
			return err
		}
		return tx.Debug().Model(&entity.Contract{}).Where(lst).Updates(&updateContract).WithContext(ctx).Error
	})
}

func (cr *contractRepo) DeleteContract(ctx context.Context, filter *models.Filter) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		var lst []entity.Contract
		lst, err := cr.Search(ctx, filter)
		if err != nil {
			return err
		}
		return tx.Debug().Model(&entity.Contract{}).Where(lst).Delete(&entity.Contract{}).WithContext(ctx).Error
	})
}

func executeSearch(filter models.Filter, db *gorm.DB, lst []entity.Contract) ([]entity.Contract, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		if filter.StudentCode != nil {
			tx = tx.Where("student_code = ?", filter.StudentCode)
		}
		if filter.Email != nil {
			tx = tx.Where("email = ?", *filter.Email)
		}
		if filter.FullName != nil {
			tx = tx.Where("full_name = ?", *filter.FullName)
		}
		if filter.Phone != nil {
			tx = tx.Where("phone = ?", *filter.Phone)
		}
		if filter.Sign != nil {
			tx = tx.Where("sign = ?", *filter.Sign)
		}
		if filter.RoomID != nil {
			tx = tx.Where("room_id = ?", *filter.RoomID)
		}
		if filter.DOB != nil {
			tx = tx.Where("dob = ?", *filter.DOB)
		}
		if filter.Gender != nil {
			tx = tx.Where("gender = ?", *filter.Gender)
		}
		if filter.IsActive != nil {
			tx = tx.Where("is_active = ?", *filter.IsActive)
		}
		if filter.NotificationChannels != nil {
			tx = tx.Where("notification_channels = ?", *filter.NotificationChannels)
		}
		if filter.Address != nil {
			tx = tx.Where("address = ?", *filter.Address)
		}
		return tx.Debug().Model(&entity.Contract{}).Find(&lst).Error
	})

	return lst, err
}

func (cr *contractRepo) Search(ctx context.Context, filter *models.Filter) ([]entity.Contract, error) {
	var lst []entity.Contract
	lst, err := executeSearch(*filter, cr.db, lst)
	if len(lst) == 0 {
		return lst, errorx.New(http.StatusNotFound, "Your contract does not exist", errors.New("contract not found"))
	}
	if err != nil {
		return lst, errorx.New(http.StatusInternalServerError, "Server error while searching contract", err)
	}
	return lst, nil
}
