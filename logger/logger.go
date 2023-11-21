package logger

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/thutgtz/go-logger/logger/constant"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/thutgtz/go-logger/logger/model"
)

type Logger interface {
	LogRequest(reqTime time.Time)
	LogResponse(reqTime time.Time)
	Info(message string)
	Error(message string)
	Debug(message string)
}

const LOGGER = "LOGGER"

func Set(ctx *fiber.Ctx) {
	logInstance := newLoggerImpl(ctx)
	context := ctx.Context()
	context.SetUserValue(LOGGER, logInstance)
}

func Get(ctx *fiber.Ctx) Logger {
	context := ctx.Context()
	return context.UserValue(LOGGER).(Logger)
}

type LoggerImpl struct {
	ctx *fiber.Ctx
}

func newLoggerImpl(ctx *fiber.Ctx) Logger {
	return &LoggerImpl{
		ctx: ctx,
	}
}

func (l *LoggerImpl) LogRequest(reqTime time.Time) {
	userId := l.ctx.GetRespHeader(string(constant.USER_ID))
	correlationId := l.ctx.GetRespHeader(string(constant.CORRELATION_ID))
	if correlationId == "" {
		correlationId = uuid.New().String()
	}

	logReq := model.RequestLogModel{
		LogType:       constant.REQUEST_LOG,
		IpAddress:     l.ctx.IP(),
		CorrelationId: correlationId,
		UserId:        userId,
		Method:        l.ctx.Method(),
		Uri:           l.ctx.Path(),
		RawUri:        l.ctx.BaseURL() + l.ctx.Path() + l.ctx.Context().QueryArgs().String(),
		ReqHeader:     string(l.ctx.Request().Header.RawHeaders()),
		ReqBody:       string(l.ctx.BodyRaw()),
		ReqTime:       reqTime.Format(time.RFC3339),
	}

	info(
		"request",
		convertStructToLogField[model.RequestLogModel](logReq)...,
	)
}

func (l *LoggerImpl) LogResponse(reqTime time.Time) {
	userId := l.ctx.GetRespHeader(string(constant.USER_ID))
	correlationId := l.ctx.GetRespHeader(string(constant.CORRELATION_ID))

	resp := model.ResponseModel{}
	json.Unmarshal(l.ctx.Response().Body(), &resp)

	execTime := time.Now().Sub(reqTime).Milliseconds()
	respLog := model.RequestLogModel{
		LogType:        constant.REQUEST_LOG,
		IpAddress:      l.ctx.IP(),
		CorrelationId:  correlationId,
		UserId:         userId,
		RespBody:       string(l.ctx.Response().Body()),
		RespHttpStatus: fmt.Sprint(l.ctx.Response().StatusCode()),
		RespStatus:     fmt.Sprint(resp.Status.Code),
		ExecTime:       fmt.Sprint(execTime),
	}

	info(
		"response",
		convertStructToLogField[model.RequestLogModel](respLog)...,
	)
}

func (l *LoggerImpl) ApiLogMetaData() model.ApiLogModel {
	userId := l.ctx.GetRespHeader(string(constant.USER_ID))
	correlationId := l.ctx.GetRespHeader(string(constant.CORRELATION_ID))

	return model.ApiLogModel{
		LogType:       constant.API_LOG,
		IpAddress:     l.ctx.IP(),
		CorrelationId: correlationId,
		UserId:        userId,
	}
}

func (l *LoggerImpl) Info(message string) {
	info(
		message,
		convertStructToLogField[model.ApiLogModel](l.ApiLogMetaData())...,
	)
}

func (l *LoggerImpl) Debug(message string) {
	debug(
		message,
		convertStructToLogField[model.ApiLogModel](l.ApiLogMetaData())...,
	)
}

func (l *LoggerImpl) Error(message string) {
	errors(
		message,
		convertStructToLogField[model.ApiLogModel](l.ApiLogMetaData())...,
	)
}
