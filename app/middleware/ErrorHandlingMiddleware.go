package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang_api/app/dtos"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				cutomError := dtos.ErrorResponse{}
				switch err.(type) {
				case dtos.ErrorResponse:
					// This section for handling from own panic/error
					cutomError = err.(dtos.ErrorResponse)
				default:
					// This section for handling unexpected error
					cutomError = dtos.ErrorResponse{
						ErrorCode: http.StatusInternalServerError,
						Message: dtos.Response{
							Status: dtos.BaseResponse{
								Success: false,
								Message: "Internal Server Error",
							},
							Data: err,
						}}
				}
				c.JSON(cutomError.ErrorCode, cutomError.Message)
				c.Abort()
			}
		}()
		c.Next()
	}
}
