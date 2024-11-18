package util

import (
	"app/domain/dto"
)

func BuildSuccessResponse[T any](data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		Data:    data,
		Message: Success.String(),
	}
}
func BuildEmptySuccessResponse() dto.ApiResponse[map[string]interface{}] {
	return dto.ApiResponse[map[string]interface{}]{
		Data:    make(map[string]interface{}),
		Message: Success.String(),
	}
}
func BuildErrorResponse(err ErrorCode) dto.ApiResponse[map[string]interface{}] {
	return dto.ApiResponse[map[string]interface{}]{
		Data:    make(map[string]interface{}),
		Message: err.String(),
		Error:   err.String(),
	}
}

func BuildErrorResponseFromError(err error) dto.ApiResponse[map[string]interface{}] {
	apiErr, ok := err.(*ApiError)
	if !ok {
		return dto.ApiResponse[map[string]interface{}]{
			Data:    make(map[string]interface{}),
			Message: Internal.String(),
			Error:   Internal.String(),
		}
	} else {
		return dto.ApiResponse[map[string]interface{}]{
			Data:    make(map[string]interface{}),
			Message: apiErr.Status.String(),
			Error:   apiErr.Status.String(),
		}
	}
}
