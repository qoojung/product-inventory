package util

import (
	"app/domain/dto"
)

func BuildSuccessResponse_[T any](data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		Data:    data,
		Message: "success",
	}
}
func BuildEmptySuccessResponse_() dto.ApiResponse[map[string]interface{}] {
	return dto.ApiResponse[map[string]interface{}]{
		Data:    make(map[string]interface{}),
		Message: "success",
	}
}
