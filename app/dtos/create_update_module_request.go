package dtos

type CreateUpdateModuleRequest struct {
	Code   string `json:"code" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Remark string `json:"remark"`
}
