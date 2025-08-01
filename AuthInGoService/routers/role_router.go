package routers

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	roleController *controllers.RoleController
}

func NewRoleRouter(_roleController *controllers.RoleController) Router {
	return &RoleRouter{
		roleController: _roleController,
	}
}

func (rr *RoleRouter) Register(r chi.Router){
	r.Get("/role/{id}" , rr.roleController.GetRoleById)
	r.Get("/role" , rr.roleController.GetRoleByName) 
	r.With(middlewares.CreateRoleRequestValidator).Post("/role" , rr.roleController.CreateRole)
	r.Delete("/role/{id}" , rr.roleController.DeleteById)
	r.With(middlewares.UpdateRoleRequestValidator).Put("/role/{id}" , rr.roleController.UpdateById)

	r.Get("/role/{id}/permissions", rr.roleController.GetRolePermissions)
	r.With(middlewares.AssignPermissionRequestValidator).Post("/role/{id}/permissions", rr.roleController.AssignPermissionToRole)
	r.With(middlewares.RemovePermissionRequestValidator).Delete("/role/{id}/permissions", rr.roleController.RemovePermissionFromRole)
	r.Get("/role-permissions", rr.roleController.GetAllRolePermissions)
	r.With(middlewares.JWTMiddleware , middlewares.RequireAllRoles("admin")).Post("/role/{userId}/assign/{roleId}", rr.roleController.AssignRoleToUser)
}