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

	result, err := ctrl.service.GetUserByID(uriUuid.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	user := dtos.CreateUserResponse{
		Id:          result.Id,
		Username:    result.Username,
		Nickname:    result.Nickname,
		Email:       result.Email,
		CreatedDate: result.CreatedDate,
		UpdatedDate: result.UpdatedDate,
		RoleId:      result.RoleId,
	}

	response := dtos.Response{
		Status: dtos.BaseResponse{true, "Successfully"},
		Data:   user,
	}

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

	result, err := ctrl.service.GetAllUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	users := []dtos.CreateUserResponse{}
	for _, user := range *result {
		users = append(users, dtos.CreateUserResponse{
			Id:          user.Id,
			Username:    user.Username,
			Nickname:    user.Nickname,
			Email:       user.Email,
			CreatedDate: user.CreatedDate,
			UpdatedDate: user.UpdatedDate,
			RoleId:      user.RoleId,
		})
	}

	response := dtos.Response{
		Status: dtos.BaseResponse{true, "Successfully"},
		Data:   users,
	}

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
		response := dtos.Response{
			Status: dtos.BaseResponse{false, "Failed On Falidation"},
			Data:   errMsg,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err := req.Validate(); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		response := dtos.Response{
			Status: dtos.BaseResponse{false, "Failed On Falidation"},
			Data:   errMsg,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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

	response := dtos.Response{
		Status: dtos.BaseResponse{true, "Successfully"},
		Data:   createdUser,
	}

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
		response := dtos.Response{
			Status: dtos.BaseResponse{false, "Failed On Validation"},
			Data:   errMsg,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err := req.Validate(); err != nil {
		errMsg := tools.GenerateErrorMessageV2(err)
		response := dtos.Response{
			Status: dtos.BaseResponse{false, "Failed On Validation"},
			Data:   errMsg,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctrl.service.UpdateUser(uriUuid.Id, req)
	if err != nil {
		response := dtos.Response{
			Status: dtos.BaseResponse{false, err.Error()},
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := dtos.Response{
		Status: dtos.BaseResponse{true, "Update Successfully"},
	}

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
		response := dtos.Response{
			Status: dtos.BaseResponse{false, "Failed On Validation"},
			Data:   errMsg,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctrl.service.DeleteUser(uriUuid.Id)
	if err != nil {
		response := dtos.Response{
			Status: dtos.BaseResponse{false, err.Error()},
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := dtos.Response{
		Status: dtos.BaseResponse{true, "Delete Successfully"},
	}

	ctx.JSON(http.StatusOK, response)
}
