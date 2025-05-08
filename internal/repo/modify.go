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
		return tx.Debug().Model(&entity.Contract{}).Create(&createContract).Error
	})
}

func (cr *contractRepo) UpdateContract(ctx context.Context, keyword any, updateContract *entity.Contract) error {

	return cr.db.Transaction(func(tx *gorm.DB) error {
		return tx.Debug().Model(&entity.Contract{}).Where("id = ?", keyword).Updates(&updateContract).Error
	})
}

func (cr *contractRepo) DeleteContract(ctx context.Context, keyword any) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		return tx.Debug().Model(&entity.Contract{}).Where("id = ?", keyword).Delete(&entity.Contract{}).Error
	})
}

func executeSearch(filter models.Filter, db *gorm.DB, lst []entity.Contract) (error, []entity.Contract) {
	err := db.Transaction(func(tx *gorm.DB) error {
		if filter.StudentCode != "" {
			tx = tx.Where("student_code = ?", filter.StudentCode)
		}
		if filter.Email != "" {
			tx = tx.Where("email = ?", filter.Email)
		}
		if filter.FullName != "" {
			tx = tx.Where("full_name = ?", filter.FullName)
		}
		if filter.Phone != "" {
			tx = tx.Where("phone = ?", filter.Phone)
		}
		if filter.Sign != "" {
			tx = tx.Where("sign = ?", filter.Sign)
		}
		if filter.RoomID != "" {
			tx = tx.Where("room_id = ?", filter.RoomID)
		}
		if !filter.DOB.IsZero() {
			tx = tx.Where("dob = ?", filter.DOB)
		}
		if filter.Gender != 0 {
			tx = tx.Where("gender = ?", filter.Gender)
		}

		return tx.Debug().Model(&entity.Contract{}).Find(&lst).Error
	})

	return err, lst
}

func (cr *contractRepo) Search(ctx context.Context, filter *models.Filter) ([]entity.Contract, error) {
	var lst []entity.Contract
	err, lst := executeSearch(*filter, cr.db, lst)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return lst, errorx.WrapError(err, http.StatusNotFound, "Contract not found")
		}
		return lst, errorx.WrapError(err, http.StatusInternalServerError, "Server error while searching contract")
	}
	return lst, nil
}
