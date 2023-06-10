package dtos

import (
	"github.com/go-playground/validator/v10"
)

type CreateOrUpdateUserRequest struct {
	Username        string `json:"username" validate:"required"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=5,eqfield=Password"`
	RoleId          string `json:"roleId" validate:"uuid,required"`
}

func (c *CreateOrUpdateUserRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}
	return nil
}
