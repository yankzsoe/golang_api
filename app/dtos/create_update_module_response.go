package dtos

import "time"

type CreateUpdateModuleResponse struct {
	Id          string
	Name        string
	Remark      string
	CreatedDate time.Time
	UpdateDate  *time.Time
}
