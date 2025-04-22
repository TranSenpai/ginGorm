package service

import (
	"errors"
	model "main/internal/models"
	repo "main/internal/repo"
)

var (
	RoomService *roomService
)

func GetRoomService() *roomService {
	if RoomService == nil {
		RoomService = &roomService{}
	}
	return RoomService
}

type roomService struct {
}

func (r roomService) CreateRoom(room *model.Room) error {
	roomRepo := repo.GetInstanceRoom()

	err := roomRepo.RegisterRoom(room)
	if err != nil {
		return errors.New("can not create new room")
	}

	return nil
}

func (r roomService) DeleteRoom(id string) error {
	roomRepo := repo.GetInstanceRoom()

	err := roomRepo.DeleteRoom(id)
	if err != nil {
		return errors.New("can not create new room")
	}

	return nil
}

func (r roomService) UpdateRoom(id string, room *model.Room) error {
	roomRepo := repo.GetInstanceRoom()

	err := roomRepo.UpdateRoom(id, room)
	if err != nil {
		return errors.New("can not update")
	}

	return nil
}

func (r roomService) Search(id string) (model.Room, error) {
	roomRepo := repo.GetInstanceRoom()

	room, err := roomRepo.Search(id)
	if err != nil {
		return model.Room{}, errors.New("can not search")
	}
	return room, nil
}

func (r roomService) SearchAll() ([]model.Room, error) {
	roomRepo := repo.GetInstanceRoom()

	room, err := roomRepo.SearchAll()
	if err != nil {
		return nil, errors.New("can not search")
	}
	return room, nil
}
