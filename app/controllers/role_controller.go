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
//	@Summary		Get Role Data
//	@Description	Get Role Data By ID
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path	string	true	"ID"
//	@Router			/role/{id} [get]
func (r *RoleController) GetRoleById(ctx *gin.Context) {
	uriId := dtos.UriUuid{}

	if err := ctx.ShouldBindUri(&uriId); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := r.service.GetRoleById(uriId)

	if result == nil {
		response := tools.CreateSuccessResponseWithData(result)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// GetRole godoc
//
//	@Summary		Get Role Data
//	@Description	Get Role Data By Name
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
		response := tools.CreateSuccessResponseWithData(result)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}

// GetRole godoc
//
//	@Summary		Get All Role Data
//	@Description	Get All Role Data
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			name	query	dtos.CommonParam	true	"Parameters"
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
		response := tools.CreateSuccessResponseWithData(result)
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)
	ctx.JSON(http.StatusOK, response)
}
