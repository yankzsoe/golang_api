package models

import "time"

type RoleModel struct {
	ID          string     `gorm:"column:role_id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string     `gorm:"column:role_name;not null;size:100"`
	IsActive    bool       `gorm:"column:is_active;not null;default:"`
	CreatedDate time.Time  `gorm:"column:created_date;autoCreateTime:true;"`
	UpdatedDate *time.Time `gorm:"column:updated_date;autoUpdateTime:false"`
}

func (c *RoleModel) TableName() string {
	return "role"
}
