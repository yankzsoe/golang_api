package tools

import (
	"golang_api/app/dtos"
)

func ThrowException(errCode int, errMessage string) {
	panic(dtos.ErrorResponse{
		ErrorCode: errCode,
		Message: dtos.Response{
			Status: dtos.BaseResponse{
				Success: false,
				Message: errMessage,
			},
		},
	})
}

func ThrowExceptionOnValidation(errCode int, data interface{}) {
	panic(dtos.ErrorResponse{
		ErrorCode: errCode,
		Message: dtos.Response{
			Status: dtos.BaseResponse{
				Success: false,
				Message: "Failed On Validation",
			},
			Data: data,
		},
	})
}
