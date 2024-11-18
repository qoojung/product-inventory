package util

import "testing"

func TestErrorCode_String(t *testing.T) {
	tests := []struct {
		name string
		e    ErrorCode
		want string
	}{
		{"Argument", Argument, "argument"},
		{"Success", Success, "success"},
		{"NotFound", NotFound, "not found"},
		{"Internal", Internal, "internal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("ErrorCode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
