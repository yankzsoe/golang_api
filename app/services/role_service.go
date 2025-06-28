package services

import (
	"golang_api/app/dtos"
	"golang_api/app/models"
	"golang_api/app/repositories"
	"golang_api/tools"
	"net/http"
	"time"
)

type RoleService struct {
	roleRepository repositories.RoleReporitory
}

func NewRoleService(role repositories.RoleReporitory) *RoleService {
	return &RoleService{
		roleRepository: role,
	}
}

func (r *RoleService) CreateRole(data dtos.CreateUpdateRoleRequest) *dtos.CreateUpdateRoleResponse {
	role := models.RoleModel{
		Code:     data.Code,
		Name:     data.Name,
		IsActive: data.IsActive,
	}

	result := r.roleRepository.Create(&role)

	response := dtos.CreateUpdateRoleResponse{
		Id:          result.ID,
		Code:        result.Code,
		Name:        result.Name,
		IsActive:    result.IsActive,
		CreatedDate: result.CreatedDate,
		UpdateDate:  result.UpdatedDate,
	}

	return &response
}

func (r *RoleService) GetRoles(param dtos.CommonParam) *[]dtos.CreateUpdateRoleResponse {
	result := r.roleRepository.FindAll(param)

	if len(*result) > 0 {
		roles := []dtos.CreateUpdateRoleResponse{}
		for _, role := range *result {
			roles = append(roles, dtos.CreateUpdateRoleResponse{
				Id:          role.ID,
				Code:        role.Code,
				Name:        role.Name,
				IsActive:    role.IsActive,
				CreatedDate: role.CreatedDate,
				UpdateDate:  role.UpdatedDate,
			})
		}
		return &roles
	}

	return nil
}

func (r *RoleService) GetRoleById(id dtos.UriUuid) *dtos.CreateUpdateRoleResponse {
	result := r.roleRepository.FindById(id)

	if len(result.ID) > 0 {
		data := dtos.CreateUpdateRoleResponse{
			Id:          result.ID,
			Code:        result.Code,
			Name:        result.Name,
			IsActive:    result.IsActive,
			CreatedDate: result.CreatedDate,
			UpdateDate:  result.UpdatedDate,
		}
		return &data
	}

	return nil
}

func (r *RoleService) GetRoleByName(name dtos.UriName) *[]dtos.CreateUpdateRoleResponse {
	result := r.roleRepository.FindByName(name.Name)

	if len(*result) > 0 {
		roles := []dtos.CreateUpdateRoleResponse{}
		for _, role := range *result {
			roles = append(roles, dtos.CreateUpdateRoleResponse{
				Id:          role.ID,
				Code:        role.Code,
				Name:        role.Name,
				IsActive:    role.IsActive,
				CreatedDate: role.CreatedDate,
				UpdateDate:  role.UpdatedDate,
			})
		}
		return &roles
	}

	return nil
}

func (r *RoleService) GetRoleWithModule(name dtos.UriName) *dtos.RoleWithModuleResponse {
	return r.roleRepository.FindRoleWithModule(name.Name)
}

func (r *RoleService) PutRole(id dtos.UriUuid, role dtos.CreateUpdateRoleRequest) *dtos.CreateUpdateRoleResponse {
	timeNow := time.Now()
	data := models.RoleModel{
		ID:          id.Id,
		Code:        role.Code,
		Name:        role.Name,
		IsActive:    role.IsActive,
		UpdatedDate: &timeNow,
	}

	result, err := r.roleRepository.Update(&data)
	if err != nil {
		tools.ThrowException(http.StatusInternalServerError, err.Error())
	}

	if len(result.ID) < 1 {
		return nil
	}

	updatedRole := dtos.CreateUpdateRoleResponse{
		Id:          result.ID,
		Code:        role.Code,
		Name:        result.Name,
		IsActive:    result.IsActive,
		CreatedDate: result.CreatedDate,
		UpdateDate:  result.UpdatedDate,
	}

	return &updatedRole
}

func (r *RoleService) PutRoleSetModule(id dtos.UriUuid, role dtos.RoleSetModuleRequest) *dtos.RoleSetModuleResponse {
	modules := []models.RoleModuleModel{}
	roleId := id.Id

	// Mapping from the dto to the model
	for _, module := range role.Modules {
		modules = append(modules, models.RoleModuleModel{
			RoleId:    roleId,
			ModuleId:  module.ModuleId,
			CanRead:   module.CanRead,
			CanCreate: module.CanCreate,
			CanUpdate: module.CanUpdate,
			CanDelete: module.CanDelete,
		})
	}

	// Process the data in the model into a repository
	result, err := r.roleRepository.UpdateSetRole(roleId, &modules)
	if err != nil {
		tools.ThrowException(http.StatusInternalServerError, err.Error())
	}

	// result := dtos.RoleSetModuleResponse{}
	// for _, module := range modules {
	// 	result.Modules = append(result.Modules, dtos.ModuleDetail{
	// 		ModuleId:  module.ID,
	// 		CanRead:   module.CanRead,
	// 		CanCreate: module.CanCreate,
	// 		CanUpdate: module.CanUpdate,
	// 		CanDelete: module.CanDelete,
	// 	})
	// }

	return result
}

func (r *RoleService) DeleteRole(id dtos.UriUuid) *dtos.CreateUpdateRoleResponse {
	result, err := r.roleRepository.Delete(id.Id)
	if err != nil {
		tools.ThrowException(http.StatusInternalServerError, err.Error())
	}

	if len(result.ID) > 1 {
		data := &dtos.CreateUpdateRoleResponse{
			Id:          result.ID,
			Code:        result.Code,
			Name:        result.Name,
			IsActive:    result.IsActive,
			CreatedDate: result.CreatedDate,
			UpdateDate:  result.UpdatedDate,
		}
		return data
	}

	return nil
}
