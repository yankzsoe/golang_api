package dtos

type RoleSetModuleRequest struct {
	Modules []RoleModule `json:"modules"`
}

type RoleModule struct {
	ModuleId  string `json:"moduleId" binding:"required"`
	CanRead   bool   `json:"canRead"`
	CanCreate bool   `json:"canCreate"`
	CanUpdate bool   `json:"canUpdate"`
	CanDelete bool   `json:"canDelete"`
}
