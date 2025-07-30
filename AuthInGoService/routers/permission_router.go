package routers

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type PermissionRouter struct {
	permissionController *controllers.PermissionController
}

func NewPermissionRouter(_permissionController *controllers.PermissionController) Router {
	return &PermissionRouter{
		permissionController: _permissionController,
	}
}

func (rr *PermissionRouter) Register(r chi.Router){
	r.Get("/permission/{id}" , rr.permissionController.GetPermissionById)
	r.Get("/permission" , rr.permissionController.GetPermissionByName)
	r.Post("/permission" , rr.permissionController.CreatePermission)
	r.Delete("/permission/{id}" , rr.permissionController.DeleteById)
	r.Put("/permission/{id}" , rr.permissionController.UpdateById)
}
