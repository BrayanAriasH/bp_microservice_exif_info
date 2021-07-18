package model

type ErrorResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code uint, err error) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}
}
