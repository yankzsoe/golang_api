package dtos

type CreateUpdateModuleRequest struct {
	Name   string `json:"name" validate:"required"`
	Remark string `json:"remark"`
}
