package dtos

type RoleWithModuleResponse struct {
	RoleId   string   `json:"roleId"`
	RoleCode string   `json:"roleCode"`
	RoleName string   `json:"roleName"`
	Modules  []Module `json:"modules"`
}

type Module struct {
	ModuleId   string `json:"moduleId" gorm:"module_id"`
	ModuleCode string `json:"moduleCode" gorm:"module_code"`
	ModuleName string `json:"moduleName" gorm:"module_name"`
}

type RoleWithModuleScanRow struct {
	RoleId     string `json:"role_id"`
	RoleCode   string `json:"role_code"`
	RoleName   string `json:"role_name"`
	ModuleId   string `json:"module_id"`
	ModuleCode string `json:"module_code"`
	ModuleName string `json:"module_name"`
}
