package controllers

import (
	"golang_api/app/dtos"
	"golang_api/app/services"
	"golang_api/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ModuleController struct {
	service *services.ModuleService
}

func NewModuleController(svc *services.ModuleService) *ModuleController {
	return &ModuleController{
		service: svc,
	}
}

// GetModule godoc
//
//	@Summary		Get Module Data
//	@Description	Get Module Data By ID
//	@Tags			Module
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path	string	true	"ID"
//	@Router			/module/{id} [get]
func (r *ModuleController) GetModuleById(ctx *gin.Context) {
	uriId := dtos.UriUuid{}

	if err := ctx.ShouldBindUri(&uriId); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.GetModuleById(uriId)

	if result == nil {
		response := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// GetModule godoc
//
//	@Summary		Get Module Data
//	@Description	Get Module Data By Name
//	@Tags			Module
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			name	path	string	true	"Name"
//	@Router			/module/name/{name} [get]
func (r *ModuleController) GetModuleByName(ctx *gin.Context) {
	uriName := dtos.UriName{}

	if err := ctx.ShouldBindUri(&uriName); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.GetModuleByName(uriName)

	if result == nil {
		response := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// GetModule godoc
//
//	@Summary		Get All Module Data
//	@Description	Get All Module Data
//	@Tags			Module
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			name	query	dtos.CommonParam	true	"Parameters"
//	@Router			/module [get]
func (r *ModuleController) GetModules(ctx *gin.Context) {
	conv := tools.Conversion{}
	limit := conv.StrToInt(ctx.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	param := dtos.CommonParam{
		Where:  ctx.Query("where"),
		Limit:  limit,
		Offset: conv.StrToInt(ctx.Query("offset")),
	}

	if err := ctx.ShouldBindUri(&param); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.GetModules(param)

	if result == nil {
		response := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// PostModule godoc
//
//	@Summary		Create Module Data
//	@Description	Create Module Data
//	@Tags			Module
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			request	body	dtos.CreateUpdateModuleRequest	true	"Body"
//	@Router			/module [post]
func (r *ModuleController) PostModule(ctx *gin.Context) {
	param := dtos.CreateUpdateModuleRequest{}

	if err := ctx.ShouldBindJSON(&param); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.CreateModule(param)

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// PutModule godoc
//
//	@Summary		Update Module Data
//	@Description	Update Module Data
//	@Tags			Module
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id		path	string							true	"Parameters"
//	@Param			request	body	dtos.CreateUpdateModuleRequest	true	"Module"
//	@Router			/module/{id} [put]
func (r *ModuleController) PutModule(ctx *gin.Context) {
	moduleId := dtos.UriUuid{}
	if err := ctx.ShouldBindUri(&moduleId); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	module := dtos.CreateUpdateModuleRequest{}
	if err := ctx.ShouldBindJSON(&module); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.PutModule(moduleId, module)

	if result == nil {
		resp := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// DeleteModule godoc
//
//	@Summary		Delete Module Data
//	@Description	Delete Module Data
//	@Tags			Module
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path	string	true	"id"
//	@Router			/module/{id} [delete]
func (r *ModuleController) DeleteModule(ctx *gin.Context) {
	param := dtos.UriUuid{}

	if err := ctx.ShouldBindUri(&param); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.DeleteModule(param)

	if result == nil {
		resp := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	response := tools.CreateSuccessDeletedResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}
