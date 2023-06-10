package dtos

type UserWithClaimsResponse struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RoleName   string `json:"role_name"`
	ModuleName string `json:"module_name"`
	CanCreate  bool   `json:"can_create"`
	CanRead    bool   `json:"can_read"`
	CanUpdate  bool   `json:"can_update"`
	CanDelete  bool   `json:"can_delete"`
}
