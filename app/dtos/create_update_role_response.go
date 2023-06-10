package dtos

import "time"

type CreateUpdateRoleResponse struct {
	Id          string
	Name        string
	IsActive    bool
	CreatedDate time.Time
	UpdateDate  *time.Time
}
