package util

import (
	"app/domain/dto"
	"reflect"
	"testing"
)

func TestBuildEmptySuccessResponse(t *testing.T) {
	want := dto.ApiResponse[map[string]interface{}]{
		Data:    make(map[string]interface{}),
		Message: "success",
		Error:   "",
	}
	t.Run("empty", func(t *testing.T) {
		if got := BuildEmptySuccessResponse(); !reflect.DeepEqual(got, want) {
			t.Errorf("BuildEmptySuccessResponse() = %v, want %v", got, want)
		}
	})

}

func TestBuildErrorResponse(t *testing.T) {
	tests := []struct {
		name string
		args ErrorCode
		want dto.ApiResponse[map[string]interface{}]
	}{
		{"Record", NotFound, dto.ApiResponse[map[string]interface{}]{Data: make(map[string]interface{}), Message: "not found", Error: "not found"}},
		{"Argument", Argument, dto.ApiResponse[map[string]interface{}]{Data: make(map[string]interface{}), Message: "argument", Error: "argument"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildErrorResponse(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildErrorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
