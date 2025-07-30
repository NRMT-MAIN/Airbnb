package dtos

type CreateRoleRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}