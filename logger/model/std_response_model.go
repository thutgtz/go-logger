package model

type Status struct {
	Code    int    `json:"code"  binding:"required"`
	Header  string `json:"header"  binding:"required"`
	Message string `json:"message"  binding:"required"`
}

type ResponseModel struct {
	Status Status      `json:"status"  binding:"required"`
	Data   interface{} `json:"data"  binding:"required"`
}
