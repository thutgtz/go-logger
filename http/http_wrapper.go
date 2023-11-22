package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/thutgtz/go-logger/logger"
	"github.com/thutgtz/go-logger/logger/constant"
	"github.com/thutgtz/go-logger/logger/model"
	responseModel "github.com/thutgtz/go-logger/response/model"

	"github.com/gofiber/fiber/v2"
)

type HttpWrapper interface {
	Get(ctx *fiber.Ctx, path string, header map[string]string, timeOutInSecond time.Duration) (map[string]interface{}, error)
	Post(ctx *fiber.Ctx, path string, header map[string]string, body interface{}, timeOutInSecond time.Duration) (map[string]interface{}, error)
	Put(ctx *fiber.Ctx, path string, header map[string]string, body interface{}, timeOutInSecond time.Duration) (map[string]interface{}, error)
	Patch(ctx *fiber.Ctx, path string, header map[string]string, body interface{}, timeOutInSecond time.Duration) (map[string]interface{}, error)
}

type httpWrapperImpl struct {
	baseUrl string
}

func NewHttpWrapper(baseUrl string) HttpWrapper {
	return &httpWrapperImpl{baseUrl: baseUrl}
}

func (h *httpWrapperImpl) httpRequest(method string, ctx *fiber.Ctx, path string, header map[string]string, body []byte, timeOutInSecond time.Duration) (map[string]interface{}, error) {
	l := logger.Get(ctx)

	defaultTimeout := time.Duration(timeOutInSecond * time.Second)
	client := &http.Client{
		Timeout: defaultTimeout,
	}

	request, errReq := http.NewRequest(method, h.baseUrl+path, bytes.NewBuffer(body))
	request.Header.Set("content-type", "application/json")
	if errReq != nil {
		l.Error(errReq.Error())
		return nil, errReq
	}

	reqTime := time.Now()
	response, errResp := client.Do(request)

	if errResp != nil {
		l.Error(errResp.Error())
		return nil, errResp
	}
	defer response.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)

	logInfo := h.logResponseInfo(ctx, reqTime, method, path, h.baseUrl+path, header, body, response)
	l.LogApi(logInfo)
	return result, nil
}

func (h *httpWrapperImpl) logResponseInfo(
	ctx *fiber.Ctx,
	reqTime time.Time,
	medthod string,
	path string,
	rawUri string,
	header map[string]string,
	body []byte,
	resp *http.Response,
) model.RequestLogModel {
	userId := ctx.Request().Header.Peek(string(constant.USER_ID))
	correlationId := ctx.Request().Header.Peek(string(constant.CORRELATION_ID))

	responseModel := responseModel.ResponseModel[interface{}]{}
	json.Unmarshal(body, &responseModel)
	headerBytes, _ := json.Marshal(header)
	execTime := time.Now().Sub(reqTime).Milliseconds()

	log := model.RequestLogModel{
		LogType:        constant.REQUEST_LOG,
		IpAddress:      ctx.IP(),
		CorrelationId:  string(correlationId),
		UserId:         string(userId),
		Method:         medthod,
		Uri:            path,
		RawUri:         rawUri,
		ReqHeader:      string(headerBytes),
		ReqBody:        string(body),
		ReqTime:        reqTime.Format(time.RFC3339),
		RespBody:       string(body),
		RespHttpStatus: resp.Status,
		RespStatus:     fmt.Sprint(responseModel.Status.Code),
		ExecTime:       fmt.Sprint(execTime),
	}

	return log
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

func (h *httpWrapperImpl) Put(ctx *fiber.Ctx, path string, header map[string]string, body interface{}, timeOutInSecond time.Duration) (map[string]interface{}, error) {
	requstBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return h.httpRequest(http.MethodPut, ctx, path, nil, requstBody, timeOutInSecond)
}

func (h *httpWrapperImpl) Patch(ctx *fiber.Ctx, path string, header map[string]string, body interface{}, timeOutInSecond time.Duration) (map[string]interface{}, error) {
	requstBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return h.httpRequest(http.MethodPatch, ctx, path, nil, requstBody, timeOutInSecond)
}
