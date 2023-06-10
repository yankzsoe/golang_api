package dtos

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Response struct {
	Status BaseResponse `json:"status"`
	Data   interface{}  `json:"data,omitempty"`
}

type ErrorResponse struct {
	ErrorCode int
	Message   Response
}
