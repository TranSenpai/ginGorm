package repo

import (
	"errors"
	model "main/internal/models"

	"gorm.io/gorm"
)

type roomRepo struct {
	db *gorm.DB
}

var (
	RoomRepo *roomRepo
)

func CreateRoomTable(db *gorm.DB) string {
	return db.AutoMigrate(&model.Room{}).Error()
}

func GetInstanceRoom() *roomRepo {
	if RoomRepo == nil {
		once.Do(func() {
			RoomRepo = &roomRepo{db: dbConn}
		})
	}
	return RoomRepo
}

func (rr roomRepo) CreateRoomTable(t *model.Room) error {
	return rr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Migrator().AutoMigrate(&t)
		if err != nil {
			return err
		}
		return nil
	})
}

func (rr roomRepo) RegisterRoom(r *model.Room) error {
	// Đặt trong 1 transaction cho đồng nhất dữ liệu giữa các table
	return rr.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Room{}).Create(&r)

		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}

func (rr roomRepo) UpdateRoom(id string, new *model.Room) error {
	return rr.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Room{}).Where("id = ?", id).Updates(new)

		if result.Error != nil {
			return errors.New("can not update transaction")
		}

		return nil
	})
}

func (rr roomRepo) DeleteRoom(id string) error {
	return rr.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.Room{}).Where("id = ?", id).Delete(&model.Room{})

		if result.Error != nil {
			return errors.New("can not delete room")
		}

		return nil
	})
}

func (rr roomRepo) Search(id string) (model.Room, error) {
	var room model.Room
	result := rr.db.Where("id = ?", id).First(&room)

	if result.Error != nil {
		return model.Room{}, errors.New("not found")
	}

	return room, nil
}

func (rr roomRepo) SearchAll() ([]model.Room, error) {
	var lst []model.Room
	result := rr.db.Find(&lst)

	if result.Error != nil {
		return nil, errors.New("can not search")
	}

	return lst, nil
}
