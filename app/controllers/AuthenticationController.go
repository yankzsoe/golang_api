package controllers

import (
	"net/http"

	"golang_api/app/dtos"
	"golang_api/app/services"
	"golang_api/app/tools"

	"github.com/gin-gonic/gin"
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
//	@Summary		Request Token
//	@Description	Request Token for Authorization or you can login with gmail from this link [https://golang-api-6ej0.onrender.com/api/v1/auth/external/google]
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

func (ctrl *AuthenticationController) Oauth2Login(ctx *gin.Context) {
	oauth2url := tools.GenerateOAuthGoogleUrl()
	ctx.Redirect(http.StatusTemporaryRedirect, oauth2url)
}

func (ctrl *AuthenticationController) Callback(ctx *gin.Context) {
	code := ctx.Query("code")
	userInfo := tools.GetUserInfo(code)

	statusCode, token, err := ctrl.service.LoginOAuthGoogle(userInfo)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(statusCode, gin.H{"data": token})
}

// RefreshToken godoc
//
//	@Summary		Refresh Token
//	@Description	refresh token to extend token's active period
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.RefreshTokenRequest	true	"body"
//	@Router			/auth/refreshToken [post]
func (ctrl *AuthenticationController) RefeshToken(ctx *gin.Context) {
	req := dtos.RefreshTokenRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	statusCode, token, err := ctrl.service.RefreshToken(req)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(statusCode, gin.H{"data": token})
}
