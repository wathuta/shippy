package dtos

import (
	"context"
	"gateway/constants"
)

type ApiResponse struct {
	Data interface{} `json:"data"`
	*constants.CustomError
}

func Ok(ctx context.Context, data interface{}) *ApiResponse {
	return &ApiResponse{
		Data:        data,
		CustomError: constants.OK(ctx),
	}
}
