package controllers

import (
	"golang_api/app/dtos"
	"golang_api/app/services"
	"golang_api/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	service *services.RoleService
}

func NewRoleController(svc *services.RoleService) *RoleController {
	return &RoleController{
		service: svc,
	}
}

// GetRole godoc
//
//	@Summary		Get Role By ID
//	@Description	Get Role By ID
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dtos.CreateUpdateRoleResponse	"Successful response"
//	@Security		ApiKeyAuth
//	@Param			id	path	string	true	"Role ID"
//	@Router			/role/{id} [get]
func (r *RoleController) GetRoleById(ctx *gin.Context) {
	uriId := dtos.UriUuid{}

	if err := ctx.ShouldBindUri(&uriId); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.GetRoleById(uriId)

	if result == nil {
		response := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// GetRole godoc
//
//	@Summary		Get Role By Name
//	@Description	Get Role By Name
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			name	path	string	true	"Name"
//	@Router			/role/name/{name} [get]
func (r *RoleController) GetRoleByName(ctx *gin.Context) {
	uriName := dtos.UriName{}

	if err := ctx.ShouldBindUri(&uriName); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.GetRoleByName(uriName)

	if result == nil {
		response := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// GetRole godoc
//
//	@Summary		Get List Role
//	@Description	Get List Role
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200		{object}	[]dtos.CreateUpdateRoleResponse	"Successful response"
//	@Param			name	query		dtos.CommonParam				true	"Parameters"
//	@Router			/role [get]
func (r *RoleController) GetRoles(ctx *gin.Context) {
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

	result := r.service.GetRoles(param)

	if result == nil {
		response := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// GetRole godoc
//
//	@Summary		Get Role With Module By Name
//	@Description	Get Role With Module By Name
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			name	path	string	true	"name"
//	@Router			/role/module/{name} [get]
func (r *RoleController) GetRoleWithModule(ctx *gin.Context) {
	uriName := dtos.UriName{}

	if err := ctx.ShouldBindUri(&uriName); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.GetRoleWithModule(uriName)

	if len(result.RoleId) < 1 {
		response := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// PostRole godoc
//
//	@Summary		Create Role Data
//	@Description	Create Role Data
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			request	body	dtos.CreateUpdateRoleRequest	true	"Body"
//	@Router			/role [post]
func (r *RoleController) PostRole(ctx *gin.Context) {
	param := dtos.CreateUpdateRoleRequest{}

	if err := ctx.ShouldBindJSON(&param); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.CreateRole(param)

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// PutRole godoc
//
//	@Summary		Update Role Data
//	@Description	Update Role Data
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id		path	string							true	"Parameters"
//	@Param			request	body	dtos.CreateUpdateRoleRequest	true	"Role"
//	@Router			/role/{id} [put]
func (r *RoleController) PutRole(ctx *gin.Context) {
	roleId := dtos.UriUuid{}
	if err := ctx.ShouldBindUri(&roleId); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	role := dtos.CreateUpdateRoleRequest{}
	if err := ctx.ShouldBindJSON(&role); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.PutRole(roleId, role)

	if result == nil {
		resp := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// PutRoleSetModule godoc
//
//	@Summary		Set Module for Role
//	@Description	Set Module for Role
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id		path	string						true	"Parameters"
//	@Param			request	body	dtos.RoleSetModuleRequest	true	"Role"
//	@Router			/role/module/set/{id} [put]
func (r *RoleController) PutRoleSetModule(ctx *gin.Context) {
	roleId := dtos.UriUuid{}

	// Validate id
	if err := ctx.ShouldBindUri(&roleId); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	// Validate body
	moudules := dtos.RoleSetModuleRequest{}
	if err := ctx.ShouldBindJSON(&moudules); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	// Process the data on servcies
	result := r.service.PutRoleSetModule(roleId, moudules)

	if result == nil {
		resp := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// DeleteRole godoc
//
//	@Summary		Delete Role Data
//	@Description	Delete Role Data
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path	string	true	"id"
//	@Router			/role/{id} [delete]
func (r *RoleController) DeleteRole(ctx *gin.Context) {
	param := dtos.UriUuid{}

	if err := ctx.ShouldBindUri(&param); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.DeleteRole(param)

	if result == nil {
		resp := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	response := tools.CreateSuccessDeletedResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}
