package dtos

import "github.com/go-playground/validator/v10"

type GetUserByID struct {
	Id string `json:"id" validation:"required,uuid"`
}

func (c *GetUserByID) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}
	return nil
}
