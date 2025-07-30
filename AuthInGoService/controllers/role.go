package controllers

import (
	"AuthInGo/dtos"
	"AuthInGo/services"
	"AuthInGo/utils"
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
	var payload *dtos.CreateRoleRequest

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

	var payload *dtos.CreateRoleRequest

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
