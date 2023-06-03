package middleware

import (
	"strings"

	"golang_api/app/configs"
	"golang_api/app/dtos"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		// hadling for unnecessary token
		if url.Path == "/api/v1/auth/requestToken" || strings.Contains(url.Path, "/swagger/") || url.Path == "/" {
			c.Next()
			return
		}

		// Get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Header 'Authorization'"})
			panic(dtos.ErrorResponse{
				ErrorCode: 401,
				Message: dtos.Response{
					Status: dtos.BaseResponse{
						Success: false,
						Message: "Missing Header 'Authorization'",
					},
				},
			})
			// c.Abort()
			return
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
			// c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			panic(dtos.ErrorResponse{
				ErrorCode: 401,
				Message: dtos.Response{
					Status: dtos.BaseResponse{
						Success: false,
						Message: err.Error(),
					},
				},
			})
			// c.Abort()
			return
		}

		claims, ok := token.Claims.(*dtos.Claims)
		if !ok || !token.Valid {
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "Not Authorize"})
			panic(dtos.ErrorResponse{
				ErrorCode: 401,
				Message: dtos.Response{
					Status: dtos.BaseResponse{
						Success: false,
						Message: "Not Authorize",
					},
				},
			})
			// c.Abort()
			return
		}

		method := c.Request.Method
		success := ClaimChecker(method, url.Path, *claims)

		if !success {
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "Can't access this resources"})
			panic(dtos.ErrorResponse{
				ErrorCode: 401,
				Message: dtos.Response{
					Status: dtos.BaseResponse{
						Success: false,
						Message: "Can't access this resources",
					},
				},
			})
			// c.Abort()
			return
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
