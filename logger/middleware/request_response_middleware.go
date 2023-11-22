package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/thutgtz/go-logger/logger"
	"github.com/thutgtz/go-logger/logger/constant"
)

func RequestResponseLogMiddleWare(c *fiber.Ctx) error {
	logger.Set(c)

	correlationId := c.GetRespHeader(string(constant.CORRELATION_ID))
	if correlationId == "" {
		correlationId = uuid.New().String()
		fmt.Printf("new correlationId : %v\n", correlationId)
		c.Request().Header.Add(string(constant.CORRELATION_ID), correlationId)
	}

	log := logger.Get(c)

	err := c.Next()

	if err == nil {
		log.LogResponse()
	}

	return err
}
