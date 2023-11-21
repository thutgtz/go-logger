package model

import "github.com/thutgtz/go-logger/logger/constant"

type ApiLogModel struct {
	LogType       constant.LogType `logSchema:"logType"`
	IpAddress     string           `logSchema:"ipAddress"`
	CorrelationId string           `logSchema:"correlationId"`
	UserId        string           `logSchema:"userId"`
	ClassName     string           `logSchema:"className"`
}
