package repositories

import (
	"golang_api/app/dtos"
	"golang_api/app/models"
	"golang_api/configs"
	"golang_api/tools"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ModuleReporitory struct {
	DB *gorm.DB
}

func NewModuleRepository() *ModuleReporitory {
	return &ModuleReporitory{
		DB: configs.GetDB(),
	}
}

func (repo *ModuleReporitory) FindAll(param dtos.CommonParam) *[]models.ModuleModel {
	module := []models.ModuleModel{}
	result := repo.DB.Where("LOWER(module_name) LIKE ?", "%"+strings.ToLower(param.Where)+"%").Limit(param.Limit).Offset(param.Offset).Find(&module)
	if result.Error != nil {
		tools.ThrowException(500, "failed to find module.")
	}

	return &module
}

func (repo *ModuleReporitory) FindById(id dtos.UriUuid) *models.ModuleModel {
	module := models.ModuleModel{}
	result := repo.DB.Where("module_id=?", id.Id).Find(&module)
	if result.Error != nil {
		tools.ThrowException(500, "failed to find module.")
	}

	return &module
}

func (repo *ModuleReporitory) FindByName(name string) *[]models.ModuleModel {
	module := []models.ModuleModel{}
	result := repo.DB.Where("LOWER(module_name) LIKE ?", "%"+strings.ToLower(name)+"%").Find(&module)
	if result.Error != nil {
		tools.ThrowException(500, "failed to find module.")
	}

	return &module
}

func (repo *ModuleReporitory) Create(model *models.ModuleModel) *models.ModuleModel {
	result := repo.DB.Create(&model)
	if result.Error != nil {
		tools.ThrowException(500, "failed to create module.")
	}

	return model
}

func (repo *ModuleReporitory) Update(data *models.ModuleModel) (*models.ModuleModel, error) {
	module := models.ModuleModel{}

	result := repo.DB.Clauses(clause.Returning{}).Model(&module).Where("module_id=?", data.ID).Updates(data)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected < 1 {
		return nil, nil
	}

	return &module, nil
}

func (repo *ModuleReporitory) Delete(id string) (*models.ModuleModel, error) {
	role := models.ModuleModel{}
	if err := repo.DB.Clauses(clause.Returning{}).Delete(&role, "module_id", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
