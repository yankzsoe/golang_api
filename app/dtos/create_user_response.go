package dtos

import "time"

type CreateUserResponse struct {
	Id          string     `json:"id"`
	Username    string     `json:"username"`
	Nickname    string     `json:"nickname"`
	Email       string     `json:"email"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate *time.Time `json:"updatedDate"`
	RoleId      string     `json:"roleId"`
}
