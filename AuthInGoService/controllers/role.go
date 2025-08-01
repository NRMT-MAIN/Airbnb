package controllers

import (
	"AuthInGo/dtos"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	roleService services.RoleService
}

func NewRoleController(_roleService services.RoleService) *RoleController {
	return &RoleController{
		roleService: _roleService,
	}
}

func (rc *RoleController) GetRoleById(w http.ResponseWriter , r *http.Request) {
	id := chi.URLParam(r , "id")

	intId , _ := strconv.Atoi(id)

	role , err := rc.roleService.GetRoleById(int64(intId))

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in fetching role by id" , http.StatusInternalServerError , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Fetched role successfully" , http.StatusOK , role)
}

func (rc *RoleController) GetRoleByName(w http.ResponseWriter , r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		roles , err := rc.roleService.GetAllRoles()

		if err != nil {
			utils.WriteErrorJsonResponse(w , "Error in fetching all roles" , http.StatusInternalServerError , err)
			return
		}
		utils.WriteSuccessJsonResponse(w , "Fetched roles successfully" , http.StatusOK , roles)
		return 
	}

	role , err := rc.roleService.GetRoleByName(name)

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in fetching role by name" , http.StatusInternalServerError , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Fetched role successfully" , http.StatusOK , role)
}

func (rc *RoleController) CreateRole(w http.ResponseWriter  , r *http.Request) {
	var payload *dtos.CreateRoleRequestDTO

	err := utils.ReadJsonBody(r , &payload)
	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in reading json" , http.StatusInternalServerError , err)
		return
	}

	role , err := rc.roleService.CreateRole(payload)

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in creating role" , http.StatusInternalServerError , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Role created successfully" , http.StatusOK , role)
}

func (rc *RoleController) DeleteById(w http.ResponseWriter , r *http.Request) {
	id := chi.URLParam(r , "id")

	intId , _ := strconv.Atoi(id)

	err := rc.roleService.DeleteRoleById(int64(intId))

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in deleting role by id" , http.StatusInternalServerError , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Role deleted successfully" , http.StatusOK , "")
}

func (rc *RoleController) UpdateById(w http.ResponseWriter , r *http.Request) {
	id := chi.URLParam(r , "id")

	var payload *dtos.CreateRoleRequestDTO

	err := utils.ReadJsonBody(r , &payload)
	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in reading json" , http.StatusInternalServerError , err)
		return
	}

	intId , _ := strconv.Atoi(id)

	role , err := rc.roleService.UpdateRoleById(int64(intId) , payload)

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in updating role by id" , http.StatusInternalServerError , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Role updated successfully" , http.StatusOK , role)
}

func (rc *RoleController) AssignRoleToUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	roleId := chi.URLParam(r, "roleId")
	if userId == "" {
		utils.WriteErrorJsonResponse(w,  "User ID is required", http.StatusBadRequest, fmt.Errorf("missing user ID"))
		return
	}
	if roleId == "" {
		utils.WriteErrorJsonResponse(w, "Role ID is required", http.StatusBadRequest, fmt.Errorf("missing role ID"))
		return
	}

	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteErrorJsonResponse(w,  "Invalid role ID", http.StatusBadRequest, err)
		return
	}

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.WriteErrorJsonResponse(w, "Invalid user ID", http.StatusBadRequest, err)
		return
	}

	err = rc.roleService.AssignRoleToUser(userIdInt, roleIdInt)
	if err != nil {
		utils.WriteErrorJsonResponse(w , "Failed to assign role to user", http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccessJsonResponse(w, "Role assigned to user successfully", http.StatusOK, nil)
}

func (rc *RoleController) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteErrorJsonResponse(w, "Role ID is required",  http.StatusBadRequest , fmt.Errorf("missing role ID"))
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteErrorJsonResponse(w, "Invalid role ID" , http.StatusBadRequest, err)
		return
	}

	rolePermissions, err := rc.roleService.GetRolePermissions(id)
	if err != nil {
		utils.WriteErrorJsonResponse(w, "Failed to fetch role permissions", http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccessJsonResponse(w, "Role permissions fetched successfully", http.StatusOK, rolePermissions)
}

func (rc *RoleController) AssignPermissionToRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteErrorJsonResponse(w, "Role ID is required", http.StatusBadRequest, fmt.Errorf("missing role ID"))
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteErrorJsonResponse(w, "Invalid role ID", http.StatusBadRequest, err)
		return
	}

	payload := r.Context().Value("payload").(dtos.AssignPermissionRequestDTO)

	rolePermission, err := rc.roleService.AddPermissionToRole(id, payload.PermissionId)
	if err != nil {
		utils.WriteErrorJsonResponse(w, "Failed to assign permission to role", http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccessJsonResponse(w, "Permission assigned to role successfully", http.StatusCreated, rolePermission)
}

func (rc *RoleController) RemovePermissionFromRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteErrorJsonResponse(w, "Role ID is required", http.StatusBadRequest, fmt.Errorf("missing role ID"))
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteErrorJsonResponse(w, "Invalid role ID", http.StatusBadRequest, err)
		return
	}

	payload := r.Context().Value("payload").(dtos.RemovePermissionRequestDTO)

	err = rc.roleService.RemovePermissionFromRole(id, payload.PermissionId)
	if err != nil {
		utils.WriteErrorJsonResponse(w, "Failed to remove permission from role", http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccessJsonResponse(w, "Permission removed from role successfully", http.StatusOK, nil)
}

func (rc *RoleController) GetAllRolePermissions(w http.ResponseWriter, r *http.Request) {
	rolePermissions, err := rc.roleService.GetAllRolePermissions()
	if err != nil {
		utils.WriteErrorJsonResponse(w, "Failed to fetch all role permissions", http.StatusInternalServerError, err)
		return
	}

	utils.WriteSuccessJsonResponse(w, "All role permissions fetched successfully", http.StatusOK, rolePermissions)
}