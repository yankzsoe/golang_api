package models

import "time"

type ModuleModel struct {
	ID          string     `gorm:"column:module_id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Code        string     `gorm:"column:module_code;index;not null;size:50"`
	Name        string     `gorm:"column:module_name;not null;index;unique;size:100"`
	Remark      string     `gotm:"column:remark;size:250"`
	CreatedDate time.Time  `gorm:"column:created_date;autoCreateTime:true;"`
	UpdatedDate *time.Time `gorm:"column:updated_date;autoUpdateTime:false"`
}

func (c *ModuleModel) TableName() string {
	return "module"
}
