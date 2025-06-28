package dtos

type UserWithClaimsResponse struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RoleCode   string `json:"role_code"`
	RoleName   string `json:"role_name"`
	ModuleCode string `json:"module_code"`
	ModuleName string `json:"module_name"`
	CanCreate  bool   `json:"can_create"`
	CanRead    bool   `json:"can_read"`
	CanUpdate  bool   `json:"can_update"`
	CanDelete  bool   `json:"can_delete"`
}
