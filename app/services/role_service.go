package services

import (
	"golang_api/app/dtos"
	"golang_api/app/models"
	"golang_api/app/repositories"
)

type RoleService struct {
	roleRepository repositories.RoleReporitory
}

func NewRoleService(role repositories.RoleReporitory) *RoleService {
	return &RoleService{
		roleRepository: role,
	}
}

func (r *RoleService) CreateRole(data dtos.CreateUpdateRoleRequest) *models.RoleModel {
	role := models.RoleModel{
		Name:     data.Name,
		IsActive: data.IsActive,
	}

	return r.roleRepository.Create(&role)
}

func (r *RoleService) GetRoles(param dtos.CommonParam) *[]dtos.CreateUpdateRoleResponse {
	result := r.roleRepository.FindAll(param)

	if len(*result) > 0 {
		roles := []dtos.CreateUpdateRoleResponse{}
		for _, role := range *result {
			roles = append(roles, dtos.CreateUpdateRoleResponse{
				Id:          role.ID,
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
