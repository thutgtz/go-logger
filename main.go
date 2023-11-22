package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/thutgtz/go-logger/http"
	"github.com/thutgtz/go-logger/logger"
	"github.com/valyala/fasthttp"
)

func main() {
	logger.Init("")
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	logger.Set(c)
	header := map[string]string{}
	// header["sss"] = "sss"
	h := http.NewHttpWrapper("http://localhost:8000")
	resp, _ := h.Get(c, string("/api/v2/simple"), header, 10)
	fmt.Println(resp)
}
