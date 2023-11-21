package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/thutgtz/go-logger/logger"

	"github.com/gofiber/fiber/v2"
)

type HttpWrapper interface {
	Get(ctx *fiber.Ctx, path string, header map[string]string, timeOutInSecond time.Duration) (map[string]interface{}, error)
	Post(ctx *fiber.Ctx, path string, header map[string]string, body interface{}, timeOutInSecond time.Duration) (map[string]interface{}, error)
}

type httpWrapperImpl struct {
	baseUrl string
}

func NewHttpWrapper(baseUrl string) HttpWrapper {
	return &httpWrapperImpl{baseUrl: baseUrl}
}

func (h *httpWrapperImpl) httpRequest(method string, ctx *fiber.Ctx, path string, header map[string]string, body []byte, timeOutInSecond time.Duration) (map[string]interface{}, error) {
	l := logger.Get(ctx)
	l.Info(path)

	defaultTimeout := time.Duration(timeOutInSecond * time.Second)
	client := &http.Client{
		Timeout: defaultTimeout,
	}

	request, errReq := http.NewRequest(method, h.baseUrl+path, bytes.NewBuffer(body))
	request.Header.Set("content-type", "application/json")
	if errReq != nil {
		return nil, errReq
	}

	response, errResp := client.Do(request)
	if errResp != nil {
		return nil, errResp
	}
	defer response.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)

	return result, nil
}

func (h *httpWrapperImpl) Get(ctx *fiber.Ctx, path string, header map[string]string, timeOutInSecond time.Duration) (map[string]interface{}, error) {
	return h.httpRequest(http.MethodGet, ctx, path, nil, nil, timeOutInSecond)
}

func (h *httpWrapperImpl) Post(ctx *fiber.Ctx, path string, header map[string]string, body interface{}, timeOutInSecond time.Duration) (map[string]interface{}, error) {
	requstBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return h.httpRequest(http.MethodPost, ctx, path, nil, requstBody, timeOutInSecond)
}
