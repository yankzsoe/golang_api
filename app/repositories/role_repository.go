package repositories

import (
	"golang_api/app/dtos"
	"golang_api/app/models"
	"golang_api/configs"
	"golang_api/tools"
	"strings"

	"gorm.io/gorm"
)

type RoleReporitory struct {
	DB *gorm.DB
}

func NewRoleRepository() *RoleReporitory {
	return &RoleReporitory{
		DB: configs.GetDB(),
	}
}

func (repo *RoleReporitory) FindAll(param dtos.CommonParam) *[]models.RoleModel {
	roles := []models.RoleModel{}
	result := repo.DB.Where("LOWER(role_name) LIKE ?", "%"+strings.ToLower(param.Where)+"%").Limit(param.Limit).Offset(param.Offset).Find(&roles)
	if result.Error != nil {
		tools.ThrowException(500, "failed to find roles.")
	}

	return &roles
}

func (repo *RoleReporitory) FindById(id dtos.UriUuid) *models.RoleModel {
	roles := models.RoleModel{}
	result := repo.DB.Where("role_id=?", id.Id).Find(&roles)
	if result.Error != nil {
		tools.ThrowException(500, "failed to find roles.")
	}

	return &roles
}

func (repo *RoleReporitory) FindByName(name string) *[]models.RoleModel {
	roles := []models.RoleModel{}
	result := repo.DB.Where("LOWER(role_name) LIKE ?", "%"+strings.ToLower(name)+"%").Find(&roles)
	if result.Error != nil {
		tools.ThrowException(500, "failed to find roles.")
	}

	return &roles
}

func (repo *RoleReporitory) Create(model *models.RoleModel) *models.RoleModel {
	result := repo.DB.Create(&model)
	if result.Error != nil {
		tools.ThrowException(500, "failed to create role.")
	}

	return model
}
