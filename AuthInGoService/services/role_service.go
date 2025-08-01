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
	CreateRole(payload *dtos.CreateRoleRequestDTO) (*models.Role, error)
	DeleteRoleById(id int64) error
	UpdateRoleById(id int64, payload *dtos.CreateRoleRequestDTO) (*models.Role, error)
	GetRolePermissions(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermissions() ([]*models.RolePermission, error)
	AssignRoleToUser(userId int64, roleId int64) error
}

type RoleServiceImpl struct {
	roleRepository           db.RoleRepository
	rolePermissionRepository db.RolePermissionRepository
	userRoleRepository       db.UserRoleRepository
}

func NewRoleService(roleRepo db.RoleRepository, rolePermissionRepo db.RolePermissionRepository, userRoleRepo db.UserRoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository:           roleRepo,
		rolePermissionRepository: rolePermissionRepo,
		userRoleRepository:       userRoleRepo,
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

func (rs *RoleServiceImpl) CreateRole(payload *dtos.CreateRoleRequestDTO) (*models.Role , error) {
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

func (rs *RoleServiceImpl) UpdateRoleById(id int64, payload *dtos.CreateRoleRequestDTO) (*models.Role, error) {
	role , err := rs.roleRepository.UpdateById(id , payload)
	if err != nil {
		fmt.Println("Error in updating role in role service" , err)
		return nil , err
	}
	return role , nil
}

func (rs *RoleServiceImpl) GetRolePermissions(roleId int64) ([]*models.RolePermission, error) {
	return rs.rolePermissionRepository.GetRolePermissionByRoleId(roleId)
}

func (rs *RoleServiceImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	return rs.rolePermissionRepository.AddPermissionToRole(roleId, permissionId)
}

func (rs *RoleServiceImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	return rs.rolePermissionRepository.RemovePermissionFromRole(roleId, permissionId)
}

func (rs *RoleServiceImpl) GetAllRolePermissions() ([]*models.RolePermission, error) {
	return rs.rolePermissionRepository.GetAllRolePermissions()
}

func (rs *RoleServiceImpl) AssignRoleToUser(userId int64, roleId int64) error {
	return rs.userRoleRepository.AssignRoleToUser(userId, roleId)
}