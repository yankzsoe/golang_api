package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang_api/app/dtos"
	"golang_api/app/services"
)

type AuthenticationController struct {
	service *services.AuthenticationService
}

func NewAuthenticationController(service *services.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		service: service,
	}
}

// RequestToken godoc
//
//	@Summary		Request Token user
//	@Description	Request Token for Authorization
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.LoginRequest	true	"body"
//	@Router			/auth/requestToken [post]
func (ctrl *AuthenticationController) Login(ctx *gin.Context) {
	req := dtos.LoginRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	statusCode, token, err := ctrl.service.Login(req)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(statusCode, gin.H{"data": token})
}
