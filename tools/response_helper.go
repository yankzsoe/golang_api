package tools

import "golang_api/app/dtos"

func CreateSuccessResponse() dtos.Response {
	return dtos.Response{
		Status: dtos.BaseResponse{
			Success: true,
			Message: "Successfully"},
	}
}

func CreateSuccessResponseWithData(data interface{}) dtos.Response {
	return dtos.Response{
		Status: dtos.BaseResponse{
			Success: true,
			Message: "Successfully"},
		Data: data,
	}
}

func CreateSuccessDeletedResponseWithData(data interface{}) dtos.Response {
	return dtos.Response{
		Status: dtos.BaseResponse{
			Success: true,
			Message: "Deleted successfully",
		},
		Data: data,
	}
}

func CreateNotFoundResponse() dtos.Response {
	return dtos.Response{
		Status: dtos.BaseResponse{
			Success: false,
			Message: "Not Found"},
	}
}
