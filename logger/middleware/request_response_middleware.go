package middleware

import (
	"time"

	"github.com/thutgtz/go-logger/logger"

	"github.com/gofiber/fiber/v2"
)

func RequestResponseLogMiddleWare(c *fiber.Ctx) error {
	logger.Set(c)
	reqTime := time.Now()

	log := logger.Get(c)
	log.LogRequest(reqTime)

	err := c.Next()
	if err != nil {
		panic(err)
	}

	log.LogResponse(reqTime)

	return err
}
