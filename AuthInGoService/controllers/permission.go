package controllers

import (
	"AuthInGo/dtos"
	"AuthInGo/services"
	"AuthInGo/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PermissionController struct {
	permissionService services.PermissionService
}

func NewPermissionController(_roleService services.PermissionService) *PermissionController {
	return &PermissionController{
		permissionService: _roleService,
	}
}

func (pc *PermissionController) GetPermissionById(w http.ResponseWriter , r *http.Request) {
	id := chi.URLParam(r , "id")

	intId , _ := strconv.Atoi(id)

	permission , err := pc.permissionService.GetPermissionById(int64(intId))

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in fetching permission by id" , http.StatusInternalServerError , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Fetched permission successfully" , http.StatusOK , permission)
}

func (pc *PermissionController) GetPermissionByName(w http.ResponseWriter , r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		roles , err := pc.permissionService.GetAllPermissions()

		if err != nil {
			utils.WriteErrorJsonResponse(w , "Error in fetching all roles" , http.StatusInternalServerError , err)
			return
		}
		utils.WriteSuccessJsonResponse(w , "Fetched roles successfully" , http.StatusOK , roles)
		return 
	}

	permission , err := pc.permissionService.GetPermissionByName(name)

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in fetching permission by name" , http.StatusInternalServerError , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Fetched permission successfully" , http.StatusOK , permission)
}

func (pc *PermissionController) CreatePermission(w http.ResponseWriter  , r *http.Request) {
	var payload *dtos.CreatePermissionRequestDTO

	err := utils.ReadJsonBody(r , &payload)
	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in reading json" , http.StatusInternalServerError , err)
		return
	}

	permission , err := pc.permissionService.CreatePermission(payload)

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in creating permission" , http.StatusInternalServerError , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Permission created successfully" , http.StatusOK , permission)
}

func (pc *PermissionController) DeleteById(w http.ResponseWriter , r *http.Request) {
	id := chi.URLParam(r , "id")

	intId , _ := strconv.Atoi(id)

	err := pc.permissionService.DeletePermissionById(int64(intId))

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in deleting permission by id" , http.StatusInternalServerError , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Permission deleted successfully" , http.StatusOK , "")
}

func (pc *PermissionController) UpdateById(w http.ResponseWriter , r *http.Request) {
	id := chi.URLParam(r , "id")

	var payload *dtos.CreatePermissionRequestDTO

	err := utils.ReadJsonBody(r , &payload)
	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in reading json" , http.StatusInternalServerError , err)
		return
	}

	intId , _ := strconv.Atoi(id)

	permission , err := pc.permissionService.UpdatePermissionById(int64(intId) , payload)

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in updating permission by id" , http.StatusInternalServerError , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Permission updated successfully" , http.StatusOK , permission)
}
