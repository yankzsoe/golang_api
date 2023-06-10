package middleware

import (
	"strings"

	"golang_api/app/dtos"
	"golang_api/configs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		// hadling for unnecessary token
		if url.Path == "/api/v1/auth/requestToken" || url.Path == "/api/v1/auth/refreshToken" || strings.Contains(url.Path, "/api/v1/auth/external") || strings.Contains(url.Path, "/swagger/") || url.Path == "/" {
			c.Next()
			return
		}

		// Get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			panic(dtos.ErrorResponse{
				ErrorCode: 401,
				Message: dtos.Response{
					Status: dtos.BaseResponse{
						Success: false,
						Message: "Missing Header 'Authorization'",
					},
				},
			})
		}

		jwtKey := []byte(configs.GetJWTConfigurationInstance().Key)
		// Verify token
		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, &dtos.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return jwtKey, nil
		})

		if err != nil {
			panic(dtos.ErrorResponse{
				ErrorCode: 401,
				Message: dtos.Response{
					Status: dtos.BaseResponse{
						Success: false,
						Message: err.Error(),
					},
				},
			})
		}

		claims, ok := token.Claims.(*dtos.Claims)
		if !ok || !token.Valid {
			panic(dtos.ErrorResponse{
				ErrorCode: 401,
				Message: dtos.Response{
					Status: dtos.BaseResponse{
						Success: false,
						Message: "Not Authorize",
					},
				},
			})
		}

		method := c.Request.Method
		success := ClaimChecker(method, url.Path, *claims)

		if !success {
			panic(dtos.ErrorResponse{
				ErrorCode: 401,
				Message: dtos.Response{
					Status: dtos.BaseResponse{
						Success: false,
						Message: "Can't access this resources",
					},
				},
			})
		}

		c.Next()
	}
}

func ClaimChecker(method string, url string, claim dtos.Claims) bool {
	permissions := claim.Role.Permissions

	module := strings.Split(url, "/")[3]
	for _, permission := range permissions {
		if strings.EqualFold(permission.Module, module) {
			switch strings.ToLower(method) {
			case "post":
				return permission.CanCreate
			case "get":
				return permission.CanRead
			case "put":
				return permission.CanUpdate
			case "delete":
				return permission.CanDelete
			default:
				return false
			}
		}
	}

	return false
}
