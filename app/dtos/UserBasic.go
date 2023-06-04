package dtos

type UserBasic struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   string `json:"roleId"`
}
