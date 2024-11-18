package dto

type ApiResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
