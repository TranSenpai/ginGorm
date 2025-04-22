package service

import (
	model "main/internal/models"
	repo "main/internal/repo"
)

type roleService struct{}

var (
	RoleService *roleService
)

func GetRoleService() *roleService {
	if RoleService == nil {
		RoleService = &roleService{}
	}
	return RoleService
}

func (r roleService) CreateRole(role *model.Role) error {
	roleRepo := repo.GetInstanceRole()

	err := roleRepo.RegisterRole(role)
	if err != nil {
		return err
	}

	return nil
}

func (r roleService) DeleteRole(id string) error {
	roleRepo := repo.GetInstanceRole()

	err := roleRepo.DeleteRole(id)
	if err != nil {
		return err
	}

	return nil
}

func (r roleService) UpdateRole(id string, role *model.Role) error {
	roleRepo := repo.GetInstanceRole()

	err := roleRepo.UpdateRole(id, role)
	if err != nil {
		return err
	}

	return nil
}

func (r roleService) Search(id string) (model.Role, error) {
	roleRepo := repo.GetInstanceRole()

	role, err := roleRepo.Search(id)
	if err != nil {
		return model.Role{}, err
	}
	return role, nil
}

func (r roleService) SearchAll() ([]model.Role, error) {
	roleRepo := repo.GetInstanceRole()

	role, err := roleRepo.SearchAll()
	if err != nil {
		return nil, err
	}
	return role, nil
}
