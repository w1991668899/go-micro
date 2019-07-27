package lib_log

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
)

// https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
// 这样定义错误码:
// 首先看可以归类到哪个HTTP Code, 把这个Code作为错误码的前三位,然后后面三位递增
// 如果无法/不想归类到HTTP Code, 请用600开头
// 500** 代表系统错误
// 600** ws proxy错误

var (
	AllErrors map[int]string = make(map[int]string)

	ErrOK               = newErr(20000, "ok")
	ErrUnsupportedEvent = newErr(10001, "Unsupported Event")
	ErrQueryData        = newErr(10002, "Query data error, please try again later")
	ErrReqChannel       = newErr(10003, "Unsupported channel")
	ErrTooManyRequests  = newErr(10004, "Operation is too frequent, please try again later")

	ErrReqJsonData       = newErr(10005, "Json format error")
	ErrQueryType         = newErr(10006, "Unsupported req")
	ErrIllegalJson       = newErr(10007, "Json format or json field type error")
	ErrParams            = newErr(10008, "Parameter is incorrect")
	ErrServerConn        = newErr(10009, "Server error, please reconnect")
	ErrInvalidSymbol     = newErr(10010, "Invalid symbol")
	ErrInvalidMarketType = newErr(10011, "Invalid market type")
	ErrInvalidPeriod     = newErr(10012, "Invalid period")
	ErrAccountApiKeyNull = newErr(10013, "Account or apiKey must not be null")
	ErrDateNull          = newErr(10014, "Date must not be null")
	ErrSymbolNull        = newErr(10015, "Symbol can not be null")
	ErrPeriodNull        = newErr(10016, "Period can not be null")
	ErrDepthTypeNull     = newErr(10017, "DepthType can not be null")
	ErrInvalidDepthType  = newErr(10018, "Invalid depth type")
	ErrInvalidJWT        = newErr(10019, "Invalid jwt")
	ErrCurrencyNull      = newErr(10020, "Currency can not be null")

	ErrNotFound            = newErr(40400, "The resource you visited does not exist")
	ErrRequestDataFormat   = newErr(40401, "Request data format is incorrect")
	ErrRequestAccountOrder = newErr(40402, "Record not found")
	ErrInternal            = newErr(50000, "Internal error")
	ErrInternalFromString  = newErr(50001, "[Should never be returned]")

	// 数据库错误(50002-50010)
	ErrDatabase = newErr(50002, "Failed to access the database")
	// 正常不应该返回
	ErrSQLError         = newErr(50003, "SQL error")
	ErrDBRecordNotFound = newErr(50004, "DB record not found")

	// redis错误(50100-50150)
	ErrParamsIsEmpty = newErr(50100, "The params is empty")
	ErrRedis         = newErr(50101, "Failed to access the redis")

	// error for exchange_proxy
	ErrUserSecretKey       = newErr(60001, "Can't not access to user's secret key")
	ErrInvalidSecretKey    = newErr(60002, "Invalid secret key")
	ErrUserAccount         = newErr(60003, "Failed to get account")
	ErrSignature           = newErr(60004, "Signature verification failed")
	ErrSignRepeatedRequest = NewErr(60005, "Duplicate signature")
	ErrPongTimeout         = NewErr(60006, "pong message timeout")
	ErrOKPingMessage       = NewErr(60007, "ping message")
	ErrUnKnown             = NewErr(60008, "unknown error")
	ErrUnQuoteRequest      = newErr(60009, "UnQuoteRequest")
	ErrLackCRID            = newErr(60010, "Lack of CRID")
	ErrServerBusy          = newErr(60011, "Request server is too busy")

	//新定义错误码时，首先看公共错误码和各个服务中是否已经定义过相似的错误信息，有则复用，切勿滥用错误码
	//通知服务错误码(60050-60070)
	ErrNotifyType = newErr(60050, "Unsupported notify type")
	ErrTemplate   = newErr(60051, "Error template")
	ErrProvider   = newErr(60052, "Error Provider")
	ErrInvalidApp = newErr(60053, "Invalid application name")

	//新定义错误码时，首先看公共错误码和各个服务中是否已经定义过相似的错误信息，有则复用，切勿滥用错误码
	//存储服务错误码(60071-60080)
	ErrBucket = newErr(60071, "Bucket does not exist")

	//服务之间调用错误码(60201-60300)
	ErrCaptcha = newErr(60201, "captcha service transfer error")
	ErrUser    = newErr(60202, "user service transfer error")
	ErrEngine  = newErr(60203, "engine service transfer error")

	//系统层面错误码(60301-63400)
	ErrSystemStopServing = newErr(60301, "System stop Serving")

	ErrEcho      = newErr(71000, "Undefined Echo error")
	ErrMicro     = newErr(72000, "Undefined Micro error")
	errorMapping map[string]*Err
)

func init() {
	errorMapping = make(map[string]*Err)
	errorMapping[gorm.ErrRecordNotFound.Error()] = ErrNotFound
	errorMapping[gorm.ErrInvalidSQL.Error()] = ErrDatabase
	errorMapping[gorm.ErrInvalidTransaction.Error()] = ErrDatabase
	errorMapping[gorm.ErrUnaddressable.Error()] = ErrDatabase
}

func newErr(code int, msg string) *Err {
	if _, ok := AllErrors[code]; ok {
		panic("Duplicated error code!!!")
	}
	AllErrors[code] = msg
	return &Err{code, msg}
}

// NewErr registers a new error type.
func NewErr(code int, msg string) *Err {
	return newErr(code, msg)
}

type Err struct {
	Code    int
	Message string
}

func (e *Err) Error() string {
	r, _ := json.Marshal(e)
	return string(r)
}

func (e Err) GetMessage() string {
	return e.Message
}

func (e Err) GetMessageByParams(params ...interface{}) string {
	return fmt.Sprintf(e.Message, params...)
}

// ErrFromString assembles an Err from string.
func ErrFromString(str string) *Err {
	var e Err
	if err := json.Unmarshal([]byte(str), &e); err != nil || e.Code < 10000 {
		return &Err{ErrInternalFromString.Code, str}
	}
	return &e
}

// ErrFromGoErr transforms the golang error object to Err.
func ErrFromGoErr(err error) *Err {
	if xgbErr, ok := errorMapping[err.Error()]; ok {
		return xgbErr
	}

	if e, ok := err.(*Err); ok {
		return e
	}
	return ErrFromString(err.Error())
}

// IsErr indicates if the error is Err.
func IsErr(err error, forkingerr *Err) bool {
	e, ok := err.(*Err)
	if !ok {
		return false
	}
	return e.Code == forkingerr.Code
}

