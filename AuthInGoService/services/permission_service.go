package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/dtos"
	"AuthInGo/models"
	"fmt"
)

type PermissionService interface {
	GetPermissionById(id int64) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)
	CreatePermission(payload *dtos.CreatePermissionRequestDTO) (*models.Permission, error)
	DeletePermissionById(id int64) error
	UpdatePermissionById(id int64, payload *dtos.CreatePermissionRequestDTO) (*models.Permission, error)
}

type PermissionServiceImpl struct {
	permissionRepository db.PermissionRepository
}

func NewPermissionService(_permissionRepository db.PermissionRepository) PermissionService {
	return &PermissionServiceImpl{
		permissionRepository : _permissionRepository, 
	}
}

func (ps *PermissionServiceImpl) GetPermissionById(id int64) (*models.Permission , error) {
	permission , err := ps.permissionRepository.GetPermissionById(id)

	if err != nil {
		fmt.Println("Error in getting permission in permission service" , err)
		return nil , err
	}
	return permission , nil 
}

func (ps *PermissionServiceImpl) GetPermissionByName(name string) (*models.Permission , error) {
	permission , err := ps.permissionRepository.GetPermissionByName(name)

	if err != nil {
		fmt.Println("Error in getting permission in permission service" , err)
		return nil , err
	}
	return permission , nil 
}

func (ps *PermissionServiceImpl) GetAllPermissions() ([]*models.Permission , error) {
	permission , err := ps.permissionRepository.GetAllPermissions()

	if err != nil {
		fmt.Println("Error in getting all permission in permission service" , err)
		return nil , err
	}
	return permission , nil 
}

func (ps *PermissionServiceImpl) CreatePermission(payload *dtos.CreatePermissionRequestDTO) (*models.Permission , error) {
	permission , err := ps.permissionRepository.CreatePermission(payload)

	if err != nil {
		fmt.Println("Error in creating permission in permission service" , err)
		return nil , err
	}
	return permission , nil 
}

func (ps *PermissionServiceImpl) DeletePermissionById(id int64) error {
	err := ps.permissionRepository.DeletePermissionById(id)
	if err != nil {
		fmt.Println("Error in deleting permission in permission service" , err)
		return err
	}
	return nil 
}

func (ps *PermissionServiceImpl) UpdatePermissionById(id int64, payload *dtos.CreatePermissionRequestDTO) (*models.Permission, error) {
	permission , err := ps.permissionRepository.UpdatePermission(id , payload)
	if err != nil {
		fmt.Println("Error in updating permission in permission service" , err)
		return nil , err
	}
	return permission , nil
}

