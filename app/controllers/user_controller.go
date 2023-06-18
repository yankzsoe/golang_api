package controllers

import (
	"net/http"

	"golang_api/app/dtos"
	"golang_api/app/models"
	"golang_api/app/services"
	"golang_api/tools"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// GetUser godoc
//
//	@Summary		Get User Data
//	@Description	Get User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path	string	false	"ID"
//	@Router			/user/{id} [get]
func (ctrl *UserController) GetUser(ctx *gin.Context) {
	uriUuid := dtos.UriUuid{}
	if err := ctx.ShouldBindUri(&uriUuid); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	result := ctrl.service.GetUserByID(uriUuid.Id)
	if result == nil || len(result.Id) == 0 {
		resp := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)

	ctx.JSON(http.StatusOK, response)
}

// GetAllUser godoc
//
//	@Summary		Get All User Data
//	@Description	Get All User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			request	query	dtos.CommonParam	false	"param"
//	@Router			/user/ [get]
func (ctrl *UserController) GetAllUser(ctx *gin.Context) {
	conv := tools.Conversion{}
	limit := conv.StrToInt(ctx.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	req := dtos.CommonParam{
		Where:  ctx.Query("where"),
		Limit:  limit,
		Offset: conv.StrToInt(ctx.Query("offset")),
	}

	result := ctrl.service.GetAllUser(req)
	if len(*result) == 0 {
		resp := tools.CreateNotFoundResponse()
		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	response := tools.CreateSuccessResponseWithData(result)

	ctx.JSON(http.StatusOK, response)
}

// PostUser godoc
//
//	@Summary		Post User Data
//	@Description	Add new fake User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			request	body	dtos.CreateOrUpdateUserRequest	true	"User"
//	@Router			/user/ [post]
func (ctrl *UserController) PostUser(ctx *gin.Context) {
	req := dtos.CreateOrUpdateUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	if err := req.Validate(); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	m := models.UserModel{
		Username: req.Username,
		Nickname: req.Nickname,
		Email:    req.Email,
		Password: req.Password,
		RoleId:   req.RoleId,
	}

	result, err := ctrl.service.CreateUser(&m)
	if err != nil {
		tools.ThrowException(http.StatusInternalServerError, err.Error())
	}

	createdUser := dtos.CreateUserResponse{
		Id:          result.Id,
		Username:    result.Username,
		Nickname:    result.Nickname,
		Email:       result.Email,
		CreatedDate: result.CreatedDate,
		UpdatedDate: result.UpdatedDate,
		RoleId:      result.RoleId,
	}

	response := tools.CreateSuccessResponseWithData(createdUser)

	ctx.JSON(http.StatusCreated, response)
}

// PutUser godoc
//
//	@Summary		Put User Data
//	@Description	Update User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id		path	string							true	"User ID"
//	@Param			request	body	dtos.CreateOrUpdateUserRequest	true	"User"
//	@Router			/user/{id} [put]
func (ctrl *UserController) PutUser(ctx *gin.Context) {
	req := dtos.CreateOrUpdateUserRequest{}
	ctx.ShouldBindJSON(&req)

	uriUuid := dtos.UriUuid{}
	if err := ctx.ShouldBindUri(&uriUuid); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	if err := req.Validate(); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	err := ctrl.service.UpdateUser(uriUuid.Id, req)
	if err != nil {
		tools.ThrowException(http.StatusInternalServerError, err.Error())
	}

	response := tools.CreateSuccessResponse()

	ctx.JSON(http.StatusOK, response)
}

// DeleteUser godoc
//
//	@Summary		Delete User Data
//	@Description	Delete User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path	string	true	"User ID"
//	@Router			/user/{id} [delete]
func (ctrl *UserController) DeleteUser(ctx *gin.Context) {
	uriUuid := dtos.UriUuid{}
	if err := ctx.ShouldBindUri(&uriUuid); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, errMsg)
	}

	err := ctrl.service.DeleteUser(uriUuid.Id)
	if err != nil {
		tools.ThrowExceptionOnValidation(http.StatusBadRequest, err.Error())
	}

	response := tools.CreateSuccessResponse()

	ctx.JSON(http.StatusOK, response)
}
