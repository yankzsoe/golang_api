package dtos

type CreateUpdateRoleRequest struct {
	Name     string `json:"name" validate:"required"`
	Code     string `json:"code" validate:"required"`
	IsActive bool   `json:"is_active"`
}
