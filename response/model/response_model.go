package model

type Status struct {
	Code        int    `json:"code"  binding:"required"`
	Header      string `json:"header"  binding:"required"`
	Description string `json:"description"  binding:"required"`
}
type ResponseModel[T interface{}] struct {
	Status Status `json:"status"  binding:"required"`
	Data   T      `json:"data"  binding:"required"`
}
