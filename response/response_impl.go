package response

import (
	"encoding/json"
	"fmt"

	"github.com/thutgtz/go-logger/response/constant"
	"github.com/thutgtz/go-logger/response/model"

	"github.com/go-http-utils/headers"
	"github.com/gofiber/fiber/v2"
)

type CustomResponseImpl struct {
	statusLookUp map[constant.Langauge]map[int]model.Status
}

func newCustomResponseImpl(statusCodeCsvByte []byte) CustomResponse {
	c := &CustomResponseImpl{
		statusLookUp: map[constant.Langauge]map[int]model.Status{},
	}
	err := json.Unmarshal(statusCodeCsvByte, &c.statusLookUp)
	if err != nil {
		panic(fmt.Sprintf("unmarshal status code error : %v", err))
	}
	return c
}

func (c *CustomResponseImpl) getResponseStatus(ctx *fiber.Ctx, code int) model.Status {
	acceptLang := string(ctx.Request().Header.Peek(headers.AcceptLanguage))
	language := constant.GetLangauge(acceptLang)
	return c.statusLookUp[language][code]
}

func (c *CustomResponseImpl) success(ctx *fiber.Ctx, data interface{}, code int) error {
	resp := model.ResponseModel[interface{}]{
		Status: c.getResponseStatus(ctx, code),
		Data:   data,
	}
	return ctx.JSON(resp)
}

func (c *CustomResponseImpl) businessException(code int, exception ...error) error {
	if len(exception) > 0 {
		return fiber.NewError(int(code), exception[0].Error())
	}
	return fiber.NewError(int(code))
}

func (c *CustomResponseImpl) httpException(err *fiber.Error, exception ...error) error {
	if len(exception) > 0 {
		return fiber.NewError(err.Code, exception[0].Error())
	}
	return fiber.NewError(err.Code)
}
