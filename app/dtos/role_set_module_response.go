package dtos

type RoleSetModuleResponse struct {
	RoleId   string         `json:"roleId"`
	RoleName string         `json:"roleName"`
	Modules  []ModuleDetail `json:"modules"`
}

type ModuleDetail struct {
	ModuleId   string `json:"moduleId"`
	ModuleName string `json:"moduleName"`
	CanRead    bool   `json:"canRead"`
	CanCreate  bool   `json:"canCreate"`
	CanUpdate  bool   `json:"canUpdate"`
	CanDelete  bool   `json:"canDelete"`
}

type RoleSetModuleScanRows struct {
	RoleId     string `json:"role_id"`
	RoleName   string `json:"role_name"`
	ModuleId   string `json:"module_id"`
	ModuleName string `json:"module_name"`
	CanRead    bool   `json:"can_read"`
	CanCreate  bool   `json:"can_create"`
	CanUpdate  bool   `json:"can_update"`
	CanDelete  bool   `json:"can_delete"`
}
