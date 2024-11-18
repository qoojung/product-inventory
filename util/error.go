package util

import (
	"fmt"
)

type ErrorCode int

const (
	Success ErrorCode = iota
	Internal
	NotFound
	Argument
)

func (e ErrorCode) String() string {
	return [...]string{"success", "internal", "not found", "argument"}[e]
}

type ApiError struct {
	Status ErrorCode
	Err    error
}

func (r *ApiError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.Status, r.Err)
}
func NewInternalError(err error) *ApiError {
	return &ApiError{
		Status: Internal,
		Err:    err,
	}
}
