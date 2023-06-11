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
