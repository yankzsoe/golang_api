package models

type RoleModuleModel struct {
	ID          string        `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	RoleId      string        `gorm:"column:role_id;type:uuid;not null;"`
	ModuleId    string        `gorm:"column:module_id;type:uuid;not null;"`
	CanCreate   bool          `gorm:"column:can_create;default:false"`
	CanRead     bool          `gorm:"column:can_read;default:false"`
	CanUpdate   bool          `gorm:"column:can_update;default:false"`
	CanDelete   bool          `gorm:"column:can_delete;default:false"`
	Remark      string        `gorm:"column:remark;size:150;"`
	ModuleModel []ModuleModel `gorm:"foreignKey:module_id"`
	RoleModel   []RoleModel   `gorm:"foreignKey:role_id"`
}

func (c *RoleModuleModel) TableName() string {
	return "role_module"
}
