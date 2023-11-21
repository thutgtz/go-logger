package model

import "github.com/thutgtz/go-logger/logger/constant"

type RequestLogModel struct {
	LogType        constant.LogType `logSchema:"logType"`
	IpAddress      string           `logSchema:"ipAddress"`
	CorrelationId  string           `logSchema:"correlationId"`
	UserId         string           `logSchema:"userId"`
	ClassName      string           `logSchema:"className"`
	Method         string           `logSchema:"method"`
	Uri            string           `logSchema:"uri"`
	RawUri         string           `logSchema:"rawUri"`
	ReqHeader      string           `logSchema:"reqHeader"`
	ReqBody        string           `logSchema:"reqBody"`
	ReqTime        string           `logSchema:"reqTime"`
	RespHttpStatus string           `logSchema:"respHttpStatus"`
	RespStatus     string           `logSchema:"respStatus"`
	HttpStatus     string           `logSchema:"httpStatus"`
	RespBody       string           `logSchema:"respBody"`
	ExecTime       string           `logSchema:"execTime"`
}
