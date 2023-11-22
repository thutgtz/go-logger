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

	correlationId := c.Request().Header.Peek(string(constant.ACCEPT_LANGUAGE))
	if correlationId == nil {
		newCorrelationId := uuid.New().String()
		fmt.Printf("new correlationId : %v\n", newCorrelationId)
		c.Request().Header.Add(string(constant.CORRELATION_ID), newCorrelationId)
	}

	log := logger.Get(c)

	err := c.Next()

	if err == nil {
		log.LogResponse()
	}

	return err
}
