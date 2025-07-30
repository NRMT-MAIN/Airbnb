package routers

import (
	"AuthInGo/controllers"

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
	r.Post("/role" , rr.roleController.CreateRole)
	r.Delete("/role/{id}" , rr.roleController.DeleteById)
	r.Put("/role/{id}" , rr.roleController.UpdateById)
}