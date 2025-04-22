package repo

import (
	"errors"
	model "main/internal/models"

	"gorm.io/gorm"
)

type contractRepo struct {
	db *gorm.DB
}

var (
	ContractRepo *contractRepo
)

func CreateContractTable(connection *gorm.DB) string {
	return connection.AutoMigrate(&model.Contract{}).Error()
}

func GetInstanceContract() *contractRepo {
	if ContractRepo == nil {
		once.Do(func() {
			ContractRepo = &contractRepo{db: dbConn}
		})
	}
	return ContractRepo
}

func (cr contractRepo) CreateContractTable(t *model.Contract) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Migrator().AutoMigrate(&t)
		if err != nil {
			return err
		}
		return nil
	})
}

func (cr contractRepo) RegisterContract(c *model.Contract) error {
	// Đặt trong 1 transaction cho đồng nhất dữ liệu giữa các table
	return cr.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Contract{}).Create(&c)

		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}

func (cr contractRepo) UpdateContract(id string, new *model.Contract) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Contract{}).Where("id = ?", id).Updates(new)

		if result.Error != nil {
			return errors.New("can not update transaction")
		}

		return nil
	})
}

func (cr contractRepo) DeleteContract(id string) error {
	return cr.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Contract{}).Where("id = ?", id).Delete(&model.Contract{})

		if result.Error != nil {
			return errors.New("can not delete contract")
		}

		return nil
	})
}

func (cr contractRepo) Search(id string) (model.Contract, error) {
	var contract model.Contract
	result := cr.db.Where("id = ?", id).First(&contract)

	if result.Error != nil {
		return model.Contract{}, errors.New("not found")
	}

	return contract, nil
}

func (cr contractRepo) SearchAll() ([]model.Contract, error) {
	var lst []model.Contract
	result := cr.db.Find(&lst)

	if result.Error != nil {
		return nil, errors.New("can not search")
	}

	return lst, nil
}
