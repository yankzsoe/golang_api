package dummy

import (
	"time"

	"golang_api/app/models"

	"gorm.io/gorm"
)

func GetDummyRole() []models.RoleModel {
	return []models.RoleModel{
		// Create role with name Admin
		{
			ID:          "dcefcb7d-cd77-42ae-8da0-a8707487cd15",
			Code:        "admin",
			Name:        "Admin",
			IsActive:    true,
			CreatedDate: time.Now(),
		},
		// Create role with name User
		{
			ID:          "77691975-c695-4afd-aa40-286f2a26857d",
			Code:        "user",
			Name:        "User",
			IsActive:    true,
			CreatedDate: time.Now(),
		},
	}
}

func GetDummyModule() []models.ModuleModel {
	return []models.ModuleModel{
		// Create module 'User' cz only table user we will manage
		{
			ID:          "81601ca7-d217-41a6-8b1c-b7fd46fb7a91",
			Code:        "user",
			Name:        "User",
			Remark:      "User Management",
			CreatedDate: time.Now(),
		},
	}
}

func GetDummyRoleModule() []models.RoleModuleModel {
	return []models.RoleModuleModel{
		// Create Admin with full access
		{
			ID:        "5be72423-2cdd-42d4-9a3a-0fb27a4163ef",
			RoleId:    "dcefcb7d-cd77-42ae-8da0-a8707487cd15",
			ModuleId:  "81601ca7-d217-41a6-8b1c-b7fd46fb7a91",
			CanCreate: true, CanRead: true, CanUpdate: true, CanDelete: true,
		},
		// Create User with read only access
		{
			ID:        "d11f25f1-5c4e-4d76-8c98-9a355ae9bd95",
			RoleId:    "77691975-c695-4afd-aa40-286f2a26857d",
			ModuleId:  "81601ca7-d217-41a6-8b1c-b7fd46fb7a91",
			CanCreate: false, CanRead: true, CanUpdate: false, CanDelete: false,
		},
	}
}

func GetDummyUser() []models.UserModel {
	return []models.UserModel{
		// Create Admin with name admin
		// and password Admin123
		{
			Username:    "admin",
			Nickname:    "admin",
			Email:       "admin@email.com",
			Password:    "Sv94syV1ksY=",
			CreatedDate: time.Now(),
			RoleId:      "dcefcb7d-cd77-42ae-8da0-a8707487cd15",
		},
		// Create User with name user
		// and password User123
		{
			Username:    "user",
			Nickname:    "user",
			Email:       "user@email.com",
			Password:    "XuhwqHp2kw==",
			CreatedDate: time.Now(),
			RoleId:      "77691975-c695-4afd-aa40-286f2a26857d",
		},
	}
}

type CreateDummyRole struct{}

func (c *CreateDummyRole) Apply(db *gorm.DB) error {
	model := models.RoleModel{}
	count := 0
	err := db.Raw("SELECT COUNT(*) FROM " + model.TableName()).Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		// Get dummy data
		roles := GetDummyRole()

		// Insert dummy
		for _, role := range roles {
			result := db.Create(&role)
			if result.Error != nil {
				return result.Error
			}
		}

	}

	return nil
}

type CreateDummyModule struct{}

func (c *CreateDummyModule) Apply(db *gorm.DB) error {
	model := models.ModuleModel{}
	count := 0
	err := db.Raw("SELECT COUNT(*) FROM " + model.TableName()).Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		// Get dummy data
		modules := GetDummyModule()

		// Insert dummy
		for _, module := range modules {
			result := db.Create(&module)
			if result.Error != nil {
				return result.Error
			}
		}

	}

	return nil
}

type CreateDummyRoleModule struct{}

func (c *CreateDummyRoleModule) Apply(db *gorm.DB) error {
	model := models.RoleModuleModel{}
	count := 0
	err := db.Raw("SELECT COUNT(*) FROM " + model.TableName()).Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		// Get dummy data
		rolemodules := GetDummyRoleModule()

		// Insert dummy
		for _, rolemodule := range rolemodules {
			result := db.Create(&rolemodule)
			if result.Error != nil {
				return result.Error
			}
		}

	}

	return nil
}

type CreateDummyUser struct{}

func (c *CreateDummyUser) Apply(db *gorm.DB) error {
	model := models.UserModel{}
	count := 0
	err := db.Raw("SELECT COUNT(*) FROM " + model.TableName()).Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		// Get dummy data
		users := GetDummyUser()

		// Insert dummy
		for _, user := range users {
			result := db.Create(&user)
			if result.Error != nil {
				return result.Error
			}
		}

	}

	return nil
}
