package tool_http

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"go-micro/golib/lib/lib_log"
	"net/http"
	"sync"
	"time"
	"github.com/labstack/echo"
)

var (
	marketResponsePool sync.Pool
)

type WapperMarketEchoGetApi func(ctx echo.Context) (string, interface{}, error)

type MarketResponse struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Topic     string          `json:"topic"` //主题名
	Timestamp int64           `json:"ts"`
	Data      json.RawMessage `json:"data"`
}

func init() {
	marketResponsePool.New = func() interface{} {
		return new(MarketResponse)
	}
}

func WapperMarketEchoGetResponse(handle WapperMarketEchoGetApi) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		topic, data, err := handle(ctx)
		if err != nil {
			return errorMarketResponseJson(ctx, topic, err)
		} else {
			return successMarketResponseJson(ctx, topic, data)
		}
	}
}

func errorMarketResponseJson(ctx echo.Context, topic string, err error) error {
	e, ok := err.(*lib_log.Err)
	if !ok {
		e = new(lib_log.Err)
		e.Code = lib_log.ErrInternal.Code
		e.Message = err.Error()
	}

	codeErr, dataBytes := genResponseData(e, nil)
	return commonMarketResponseJsonBytes(ctx, codeErr, topic, dataBytes)
}

func successMarketResponseJson(ctx echo.Context, topic string, data interface{}) error {
	codeErr, dataBytes := genResponseData(lib_log.ErrOK, data)
	return commonMarketResponseJsonBytes(ctx, codeErr, topic, dataBytes)
}

func GenHttpMarketErrorResponse(topic string, err error) []byte {
	e, ok := err.(*lib_log.Err)
	if !ok {
		e = new(lib_log.Err)
		if result := json.Unmarshal([]byte(err.Error()), e); result != nil || e.Code < 100000 {
			e = &lib_log.Err{
				Code:    lib_log.ErrInternal.Code,
				Message: err.Error(),
			}
		}
	}

	response := &MarketResponse{
		Code:      e.Code,
		Topic:     topic,
		Timestamp: time.Now().Unix(),
		Data:      []byte("[]"),
	}
	b, _ := json.Marshal(response)
	return b
}

func GenHttpMarketSuccessResponse(topic string, data []byte) []byte {
	resp := &MarketResponse{
		Code:      lib_log.ErrOK.Code,
		Topic:     topic,
		Timestamp: time.Now().Unix(),
		Data:      data,
	}
	resData, err := json.Marshal(resp)
	if err != nil {
		logrus.Errorln(err)
	}
	return resData
}

func commonMarketResponseJsonBytes(ctx echo.Context, codeErr *lib_log.Err, topic string, dataBytes []byte) error {
	response := marketResponsePool.Get().(*MarketResponse)
	response.Code = codeErr.Code
	response.Message = codeErr.Message
	response.Topic = topic
	response.Timestamp = time.Now().Unix()
	response.Data = dataBytes
	defer marketResponsePool.Put(response)
	return ctx.JSON(http.StatusOK, response)
}
