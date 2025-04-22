package repo

import (
	"errors"
	model "main/internal/models"
	"sync"

	"gorm.io/gorm"
)

type roleRepo struct {
	db *gorm.DB
}

var (
	RoleRepo *roleRepo
	once     sync.Once
)

func CreateRoleTable(db *gorm.DB) string {
	return db.AutoMigrate(&model.Role{}).Error()
}

func GetInstanceRole() *roleRepo {
	if RoleRepo == nil {
		once.Do(func() {
			RoleRepo = &roleRepo{db: dbConn}
		})
	}
	return RoleRepo
}

func (rl roleRepo) CreateRoleTable(t *model.Role) error {
	return rl.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Migrator().AutoMigrate(&t)
		if err != nil {
			return err
		}
		return nil
	})
}

func (rl roleRepo) RegisterRole(r *model.Role) error {
	// Đặt trong 1 transaction cho đồng nhất dữ liệu giữa các table
	return rl.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Role{}).Create(&r)

		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}

func (rl roleRepo) UpdateRole(id string, new *model.Role) error {
	return rl.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Role{}).Where("id = ?", id).Updates(new)

		if result.Error != nil {
			return errors.New("can not update role")
		}

		return nil
	})
}

func (rl roleRepo) DeleteRole(id string) error {
	return rl.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Role{}).Where("id = ?", id).Delete(&model.Role{})

		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}

func (rl roleRepo) Search(id string) (model.Role, error) {
	var role model.Role
	result := rl.db.Where("id = ?", id).First(&role)

	if result.Error != nil {
		return model.Role{}, errors.New("not found")
	}

	return role, nil
}

func (rl roleRepo) SearchAll() ([]model.Role, error) {
	var lst []model.Role
	result := rl.db.Find(&lst)

	if result.Error != nil {
		return nil, errors.New("can not search")
	}

	if len(lst) == 0 {
		return nil, errors.New("dont have any role in table")
	}

	return lst, nil
}
