package tool_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"go-micro/golib/i18n"
	"go-micro/golib/lib/lib_log"
	"go-micro/golib/lib/lib_middleware/opentracing"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	responsePool sync.Pool
	pbMarshaler  jsonpb.Marshaler
)

type WapperEchoGetApi func(ctx echo.Context) (interface{}, error)
type EchoApiWrapper func(ctx echo.Context) (interface{}, error)

type Response struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

var microErrMap = map[int]string{
	400: "micro bad request",
	401: "micro unauthorized",
	403: "micro forbidden",
	404: "micro not found",
	409: "micro conflict",
	500: "micro internal server error",
}

func init() {
	responsePool.New = func() interface{} {
		return new(Response)
	}
	pbMarshaler = jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
		EnumsAsInts:  true,
	}
}

//func WapperEchoGetResponse(handle WapperEchoGetApi) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		if data, err := handle(ctx); err != nil {
//			return errorResponseJson(ctx, err)
//		} else {
//			return successResponseJson(ctx, data)
//		}
//	}
//}

func ServerErrorHandler(err error, c echo.Context) {
	code := lib_log.ErrInternal.Code
	msg := lib_log.ErrInternal.Message
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code + lib_log.ErrEcho.Code
		msg = fmt.Sprintf("%v", he.Message)
	} else {
		msg = fmt.Sprintf("%s (original message: %s)", msg, err.Error())
	}
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			c.NoContent(http.StatusOK)
		} else {
			codeErr, dataBytes := genResponseData(&lib_log.Err{
			Code:    code,
			Message: msg,
		}, nil)
			commonResponseJsonBytes(c, codeErr, dataBytes)
		}
	}
}

func successResponseJson(ctx echo.Context, data interface{}) error {
	codeErr, dataBytes := genResponseData(lib_log.ErrOK, data)
	return commonResponseJsonBytes(ctx, codeErr, dataBytes)
}

func genResponseData(codeErr *lib_log.Err, data interface{}) (*lib_log.Err, []byte) {
	if nil == data {
		data = make(map[string]interface{})
	}

	bufferObj 	:= new(bytes.Buffer)
	encoder 	:= json.NewEncoder(bufferObj)
	err 		:= encoder.Encode(data)
	dataBytes 	:= bufferObj.Bytes()

	if err != nil {
		codeErr = lib_log.ErrIllegalJson
	}
	if codeErr != lib_log.ErrOK {
		dataBytes = []byte("{}")
	}

	return codeErr, dataBytes
}

func commonResponseJsonBytes(ctx echo.Context, codeErr *lib_log.Err, dataBytes []byte) error {
	response := responsePool.Get().(*Response)
	response.Code = codeErr.Code
	response.Message = i18n.LocalizerLanguage(codeErr.Message, ctx.Request().Header.Get("Accept-Language"))
	response.Data = dataBytes
	defer responsePool.Put(response)
	return ctx.JSON(http.StatusOK, response)
}

// change tool_log
// echo api 响应, 不区分http method
// 添加micro错误处理
// 成功响应包装，适用proto3
func EchoResponseWrapper(handle EchoApiWrapper) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		data, err := handle(ctx)
		if err != nil {
			return errorResponseJson(ctx, err)
		} else {
			return successResponseJson2(ctx, data)
		}
	}
}

func errorResponseJson(ctx echo.Context, err error) error {
	buf, _ := ioutil.ReadAll(ctx.Request().Body)
	logrus.Fatal(logrus.Fields{
		"trace_id": ctx.Get(opentracing.TraceId),
		"api":      ctx.Request().RequestURI,
		"body":     string(buf),
		"user_id":  ctx.Get("user_id"),
		"method":   ctx.Request().Method,
		"err":      err,
	}, "http request err")

	e := lib_log.ErrFromGoErr(err)
	if e.Code == lib_log.ErrInternalFromString.Code { //这种情况是micro返回的错误
		e.Code = lib_log.ErrInternal.Code
		e.Message = err.Error()

		microErr := &struct {
			ID     string `json:"id"`
			Code   int    `json:"code"`
			Detail string `json:"detail"`
			Status string `json:"status"`
		}{}
		if err := json.Unmarshal([]byte(e.Message), microErr); err != nil {
			logrus.Fatal(logrus.Fields{
			"err": err,
			"api": ctx.Request().URL.Path,
		}, "fail to unmarshal micro-error")
		} else {
			e.Code = microErr.Code + lib_log.ErrMicro.Code
			// TODO
			if translated, ok := microErrMap[microErr.Code]; ok {
			e.Message = translated
		} else {
			e.Message = microErr.Detail
				logrus.Fatal(logrus.Fields{
			"err": microErr,
			}, "gateway fail to translate error message")
		}
	}
}

codeErr, dataBytes := genResponseData(e, nil)
	return commonResponseJsonBytes(ctx, codeErr, dataBytes)
}

func ErrorResponseJson(ctx echo.Context, err error) error {
	return errorResponseJson(ctx, err)
}

func successResponseJson2(ctx echo.Context, data interface{}) error {
	codeErr, dataBytes := genResponseDataByJsonpb(lib_log.ErrOK, data)
	return commonResponseJsonBytes(ctx, codeErr, dataBytes)
}

func genResponseDataByJsonpb(codeErr *lib_log.Err, data interface{}) (*lib_log.Err, []byte) {
	if nil == data {
		data = make(map[string]interface{})
	}

	proto, ok := data.(proto.Message)
	if !ok {
		return genResponseData(codeErr, data)
	}

	dataBytes := []byte("{}")
	bufferObj := new(bytes.Buffer)
	err := pbMarshaler.Marshal(bufferObj, proto)
	if err != nil {
		codeErr = lib_log.ErrIllegalJson
	} else {
		dataBytes = bufferObj.Bytes()
	}

	return codeErr, dataBytes
}
