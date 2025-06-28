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

func (repo *RoleReporitory) FindRoleWithModule(name string) *dtos.RoleWithModuleResponse {
	queryResult := dtos.RoleWithModuleResponse{}
	rows, err := repo.DB.Raw("SELECT r.role_id, r.role_code, r.role_name, mm.module_code, mm.module_id, mm.module_name"+
		" FROM \"role\" r"+
		" LEFT JOIN \"role_module\" rm ON r.role_id = rm.role_id"+
		" LEFT JOIN \"module\" mm ON rm.module_id = mm.module_id"+
		" WHERE LOWER(r.role_name) = ?", strings.ToLower(name)).Rows()
	if err != nil {
		tools.ThrowException(500, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		data := dtos.RoleWithModuleScanRow{}
		err := repo.DB.ScanRows(rows, &data)
		if err != nil {
			tools.ThrowException(500, err.Error())
		}

		if len(queryResult.Modules) == 0 {
			queryResult.RoleId = data.RoleId
			queryResult.RoleCode = data.RoleCode
			queryResult.RoleName = data.RoleName
			if len(data.ModuleId) > 0 {
				queryResult.Modules = append(queryResult.Modules, dtos.Module{
					ModuleId:   data.ModuleId,
					ModuleCode: data.ModuleCode,
					ModuleName: data.ModuleName,
				})
			}
		} else {
			queryResult.Modules = append(queryResult.Modules, dtos.Module{
				ModuleId:   data.ModuleId,
				ModuleCode: data.ModuleCode,
				ModuleName: data.ModuleName,
			})
		}
	}

	return &queryResult
}

func (repo *RoleReporitory) Create(model *models.RoleModel) *models.RoleModel {
	result := repo.DB.Create(&model)
	if result.Error != nil {
		tools.ThrowException(500, "failed to create role.")
	}

	return model
}

func (repo *RoleReporitory) Update(data *models.RoleModel) (*models.RoleModel, error) {
	role := models.RoleModel{}

	if err := repo.DB.Clauses(clause.Returning{}).Model(&role).Where("role_id=?", data.ID).Updates(map[string]interface{}{
		"role_name":    data.Name,
		"is_active":    data.IsActive,
		"updated_date": data.UpdatedDate,
	}).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (repo *RoleReporitory) UpdateSetRole(roleId string, data *[]models.RoleModuleModel) (*dtos.RoleSetModuleResponse, error) {
	result := dtos.RoleSetModuleResponse{}

	// Process using transactions
	err := repo.DB.Transaction(func(tx *gorm.DB) error {
		// Step 1, cleanup the role module data
		if err := tx.Delete(models.RoleModuleModel{}, "role_id", roleId).Error; err != nil {
			return err
		}

		// Step 2, If the role module doesn't send, the process will remove the role module only
		if len(*data) < 1 {
			return nil
		}

		// Step 3, Insert the role module data
		if err := tx.Create(&data).Error; err != nil {
			return err
		}

		// Step 4, select recently inserted the data
		rows, err := tx.Raw("SELECT r.role_id, r.role_name, mm.module_id, mm.module_name, rm.can_read, rm.can_create, rm.can_update, rm.can_delete"+
			" FROM \"role\" r"+
			" LEFT JOIN \"role_module\" rm ON r.role_id = rm.role_id"+
			" LEFT JOIN \"module\" mm ON rm.module_id = mm.module_id"+
			" WHERE r.role_id = ?", roleId).Rows()
		if err != nil {
			return err
		}

		defer rows.Close()
		isFrist := true
		for rows.Next() {
			scanRow := dtos.RoleSetModuleScanRows{}
			err := repo.DB.ScanRows(rows, &scanRow)
			if err != nil {
				return err
			}

			if isFrist {
				result.RoleId = scanRow.RoleId
				result.RoleName = scanRow.RoleName
			}

			result.Modules = append(result.Modules, dtos.ModuleDetail{
				ModuleId:   scanRow.ModuleId,
				ModuleName: scanRow.ModuleName,
				CanRead:    scanRow.CanRead,
				CanCreate:  scanRow.CanCreate,
				CanUpdate:  scanRow.CanUpdate,
				CanDelete:  scanRow.CanDelete,
			})

		}

		return nil
	})

	return &result, err
}

func (repo *RoleReporitory) Delete(id string) (*models.RoleModel, error) {
	role := models.RoleModel{}
	if err := repo.DB.Clauses(clause.Returning{}).Delete(&role, "role_id", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
