package logger

import (
	"reflect"

	"github.com/thutgtz/go-logger/logger/model"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Init(appName string) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.InitialFields = map[string]interface{}{
		"appName": appName,
	}

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func convertStructToLogField[T model.ApiLogModel | model.RequestLogModel](logStruct T) []zap.Field {
	structElem := reflect.ValueOf(logStruct)
	logField := []zap.Field{}

	for i := 0; i < structElem.NumField(); i++ {
		field := structElem.Type().Field(i)
		key := field.Tag.Get("logSchema")
		logField = append(logField, zap.String(key, structElem.Field(i).String()))
	}

	return logField
}

func info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func errors(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
