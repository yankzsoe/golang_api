package dtos

import "github.com/go-playground/validator/v10"

type GetAllUser struct {
	Id string `json:"id" validation:"required,uuid"`
}

func (c *GetAllUser) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}
	return nil
}
