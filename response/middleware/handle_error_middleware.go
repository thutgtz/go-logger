package middleware

import (
	"errors"

	"github.com/thutgtz/go-logger/logger"
	"github.com/thutgtz/go-logger/response"
	"github.com/thutgtz/go-logger/response/model"

	"github.com/gofiber/fiber/v2"
)

func HandleErrorMiddleware(ctx *fiber.Ctx, err error) error {
	l := logger.Get(ctx)

	l.Error(err.Error())
	l.LogResponse()

	var e *fiber.Error
	if errors.As(err, &e) {
		if e.Code > 1000 {
			resp := model.ResponseModel[interface{}]{
				Status: response.GetResponseStatus(ctx, (e.Code)),
			}
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(resp)
		} else {
			return ctx.Status(e.Code).SendString(fiber.NewError(e.Code).Error())
		}
	}
	return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
}
