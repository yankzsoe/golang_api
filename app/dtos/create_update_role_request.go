package dtos

type CreateUpdateRoleRequest struct {
	Name     string `json:"name" validate:"required"`
	IsActive bool   `json:"is_active"`
}
