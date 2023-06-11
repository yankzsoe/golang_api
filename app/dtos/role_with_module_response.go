package dtos

type RoleWithModuleResponse struct {
	RoleId   string   `json:"roleId"`
	RoleName string   `json:"roleName"`
	Modules  []Module `json:"modules"`
}

type Module struct {
	ModuleId   string `json:"moduleId" gorm:"module_id"`
	ModuleName string `json:"moduleName" gorm:"module_name"`
}

type RoleWithModuleScanRow struct {
	RoleId     string `json:"role_id"`
	RoleName   string `json:"role_name"`
	ModuleId   string `json:"module_id"`
	ModuleName string `json:"module_name"`
}
