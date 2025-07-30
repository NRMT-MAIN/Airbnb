package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/dtos"
	"AuthInGo/models"
	"fmt"
)

type RoleService interface {
	GetRoleById(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(payload *dtos.CreateRoleRequest) (*models.Role, error)
	DeleteRoleById(id int64) error
	UpdateRoleById(id int64, payload *dtos.CreateRoleRequest) (*models.Role, error)
}

type RoleServiceImpl struct {
	roleRepository db.RoleRepository
}

func NewRoleService(_roleRepository db.RoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository : _roleRepository , 
	}
}

func (rs *RoleServiceImpl) GetRoleById(id int64) (*models.Role , error) {
	role , err := rs.roleRepository.GetById(id)

	if err != nil {
		fmt.Println("Error in getting role in role service" , err)
		return nil , err
	}
	return role , nil 
}

func (rs *RoleServiceImpl) GetRoleByName(name string) (*models.Role , error) {
	role , err := rs.roleRepository.GetByName(name)

	if err != nil {
		fmt.Println("Error in getting role in role service" , err)
		return nil , err
	}
	return role , nil 
}

func (rs *RoleServiceImpl) GetAllRoles() ([]*models.Role , error) {
	role , err := rs.roleRepository.GetAll()

	if err != nil {
		fmt.Println("Error in getting all role in role service" , err)
		return nil , err
	}
	return role , nil 
}

func (rs *RoleServiceImpl) CreateRole(payload *dtos.CreateRoleRequest) (*models.Role , error) {
	role , err := rs.roleRepository.Create(payload)

	if err != nil {
		fmt.Println("Error in creating role in role service" , err)
		return nil , err
	}
	return role , nil 
}

func (rs *RoleServiceImpl) DeleteRoleById(id int64) error {
	err := rs.roleRepository.DeleteById(id)
	if err != nil {
		fmt.Println("Error in deleting role in role service" , err)
		return err
	}
	return nil 
}

func (rs *RoleServiceImpl) UpdateRoleById(id int64, payload *dtos.CreateRoleRequest) (*models.Role, error) {
	role , err := rs.roleRepository.UpdateById(id , payload)
	if err != nil {
		fmt.Println("Error in updating role in role service" , err)
		return nil , err
	}
	return role , nil
}

