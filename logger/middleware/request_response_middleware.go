package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/thutgtz/go-logger/logger"
	"github.com/thutgtz/go-logger/logger/constant"
)

func RequestResponseLogMiddleWare(c *fiber.Ctx) error {
	logger.Set(c)
	reqTime := time.Now()
	correlationId := c.GetRespHeader(string(constant.CORRELATION_ID))
	if correlationId == "" {
		fmt.Printf("new correlationId : %v\n", correlationId)
		correlationId = uuid.New().String()
		c.Request().Header.Add(string(constant.CORRELATION_ID), correlationId)
	}
	log := logger.Get(c)

	err := c.Next()

	log.LogResponse(reqTime)

	return err
}
