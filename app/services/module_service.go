package services

import (
	"golang_api/app/dtos"
	"golang_api/app/models"
	"golang_api/app/repositories"
	"golang_api/tools"
	"net/http"
	"time"
)

type ModuleService struct {
	moduleRepository repositories.ModuleReporitory
}

func NewModuleService(module repositories.ModuleReporitory) *ModuleService {
	return &ModuleService{
		moduleRepository: module,
	}
}

func (r *ModuleService) CreateModule(data dtos.CreateUpdateModuleRequest) *dtos.CreateUpdateModuleResponse {
	module := models.ModuleModel{
		Code:   data.Code,
		Name:   data.Name,
		Remark: data.Remark,
	}

	result := r.moduleRepository.Create(&module)

	response := dtos.CreateUpdateModuleResponse{
		Id:          result.ID,
		Code:        result.Code,
		Name:        result.Name,
		Remark:      result.Remark,
		CreatedDate: result.CreatedDate,
		UpdateDate:  result.UpdatedDate,
	}

	return &response
}

func (r *ModuleService) GetModules(param dtos.CommonParam) *[]dtos.CreateUpdateModuleResponse {
	result := r.moduleRepository.FindAll(param)

	if len(*result) > 0 {
		modules := []dtos.CreateUpdateModuleResponse{}
		for _, module := range *result {
			modules = append(modules, dtos.CreateUpdateModuleResponse{
				Id:          module.ID,
				Code:        module.Code,
				Name:        module.Name,
				Remark:      module.Remark,
				CreatedDate: module.CreatedDate,
				UpdateDate:  module.UpdatedDate,
			})
		}
		return &modules
	}

	return nil
}

func (r *ModuleService) GetModuleById(id dtos.UriUuid) *dtos.CreateUpdateModuleResponse {
	result := r.moduleRepository.FindById(id)

	if len(result.ID) > 0 {
		data := dtos.CreateUpdateModuleResponse{
			Id:          result.ID,
			Code:        result.Code,
			Name:        result.Name,
			Remark:      result.Remark,
			CreatedDate: result.CreatedDate,
			UpdateDate:  result.UpdatedDate,
		}
		return &data
	}

	return nil
}

func (r *ModuleService) GetModuleByName(name dtos.UriName) *[]dtos.CreateUpdateModuleResponse {
	result := r.moduleRepository.FindByName(name.Name)

	if len(*result) > 0 {
		modules := []dtos.CreateUpdateModuleResponse{}
		for _, module := range *result {
			modules = append(modules, dtos.CreateUpdateModuleResponse{
				Id:          module.ID,
				Code:        module.Code,
				Name:        module.Name,
				Remark:      module.Remark,
				CreatedDate: module.CreatedDate,
				UpdateDate:  module.UpdatedDate,
			})
		}
		return &modules
	}

	return nil
}

func (r *ModuleService) PutModule(id dtos.UriUuid, module dtos.CreateUpdateModuleRequest) *dtos.CreateUpdateModuleResponse {
	timeNow := time.Now()
	data := models.ModuleModel{
		ID:          id.Id,
		Code:        module.Code,
		Name:        module.Name,
		Remark:      module.Remark,
		UpdatedDate: &timeNow,
	}

	result, err := r.moduleRepository.Update(&data)
	if err != nil {
		tools.ThrowException(http.StatusInternalServerError, err.Error())
	}

	if result == nil {
		return nil
	}

	updatedModule := dtos.CreateUpdateModuleResponse{
		Id:          result.ID,
		Code:        result.Code,
		Name:        result.Name,
		Remark:      result.Remark,
		CreatedDate: result.CreatedDate,
		UpdateDate:  result.UpdatedDate,
	}

	return &updatedModule
}

func (r *ModuleService) DeleteModule(id dtos.UriUuid) *dtos.CreateUpdateModuleResponse {
	result, err := r.moduleRepository.Delete(id.Id)
	if err != nil {
		tools.ThrowException(http.StatusInternalServerError, err.Error())
	}

	if len(result.ID) > 1 {
		data := &dtos.CreateUpdateModuleResponse{
			Id:          result.ID,
			Code:        result.Code,
			Name:        result.Name,
			Remark:      result.Remark,
			CreatedDate: result.CreatedDate,
			UpdateDate:  result.UpdatedDate,
		}
		return data
	}

	return nil
}
