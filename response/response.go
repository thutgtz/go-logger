package response

import (
	"fmt"
	"os"

	"github.com/thutgtz/go-logger/response/model"

	"github.com/gofiber/fiber/v2"
)

var instance CustomResponse

type CustomResponse interface {
	getResponseStatus(ctx *fiber.Ctx, code int) model.Status
	success(ctx *fiber.Ctx, data interface{}, code int) error
	businessException(code int, exception ...error) error
	httpException(err *fiber.Error, exception ...error) error
}

func Init(pathToCsv string) {
	bytes, err := os.ReadFile(pathToCsv)
	if err != nil {
		panic(fmt.Sprintf("Load status code from csv error : %v", err))
	}
	instance = newCustomResponseImpl(bytes)
}

func Success(ctx *fiber.Ctx, data interface{}, code int) error {
	return instance.success(ctx, data, code)
}

func BusinessException(code int, exception ...error) error {
	return instance.businessException(code, exception...)
}

func HttpException(err *fiber.Error, exception ...error) error {
	return instance.httpException(err, exception...)
}

func GetResponseStatus(ctx *fiber.Ctx, code int) model.Status {
	return instance.getResponseStatus(ctx, code)
}
