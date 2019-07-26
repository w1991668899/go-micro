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

	//新定义错误码时，首先看公共错误码和各个服务中是否已经定义过相似的错误信息，有则复用，切勿滥用错误码
	//登录服务错误码(60081-60100)
	ErrInvalidUsername    = newErr(60081, "invalid username, username must be compliance mobile phone or email")
	ErrUsernameNotExist   = newErr(60082, "Username is not exist, please register first before login")
	ErrPasswordNotCorrect = newErr(60083, "user password is not correct")
	ErrUserPermanentLock  = newErr(60084, "user already permanent lock")
	ErrUserTemporaryLock  = newErr(60085, "user already temporary lock")
	ErrGoogleAuth         = newErr(60086, "google auth error, auth code is not correct")
	ErrImageAuth          = newErr(60087, "image auth error, auth code is not correct")
	ErrTokenNull          = newErr(60088, "token can not be null")
	ErrTokenInValid       = newErr(60089, "token is invalid")
	ErrDeviceNull         = newErr(60090, "device can not be null")
	ErrPwdNeedImageAuth   = newErr(60091, "user password is not correct, user next login need image auth")
	ErrImageAuthNeed      = newErr(60092, "user need image auth")
	ErrPermanentLock      = newErr(60093, "user is permanent lock")
	ErrTemporarayLock     = newErr(60094, "user is temporaray lock")
	ErrGoogleAuthNeed     = newErr(60095, "user need google auth")
	ErrUserLock           = newErr(60096, "user is lock")
	ErrDuplicatePayment   = newErr(60097, "Duplicate payment")

	//新定义错误码时，首先看公共错误码和各个服务中是否已经定义过相似的错误信息，有则复用，切勿滥用错误码
	//用户服务错误码(60101-60140)
	ErrUsernameExist               = newErr(60101, "Username already exists")
	ErrInvalidMobileOrEmail        = newErr(60102, "Please enter the correct mobile number or email address")
	ErrInvalidCaptcha              = newErr(60103, "Incorrect verification code")
	ErrAccountNotExist             = newErr(60104, "Account doesn't exist")
	ErrGoogleBound                 = newErr(60105, "Google verification has been bound")
	ErrOriginalLoginPassword       = newErr(60106, "Original password verification failed")
	ErrVerifyToken                 = newErr(60107, "The token of verification is incorrect")
	ErrOriginalFundPassword        = newErr(60108, "Original fund password verification failed")
	ErrPasswordNULL                = newErr(60109, "Password can not be null")
	ErrNoKyc                       = newErr(60110, "No kyc")
	ErrEmailAuthCode               = newErr(60111, "Invalid email verification code")
	ErrMobileAuthCode              = newErr(60112, "Invalid mobile verification code")
	ErrFundPassword                = newErr(60113, "Invalid fund password")
	ErrMobileAreaCode              = newErr(60114, "Invalid mobile area code")
	ErrLoginPasswordLen            = newErr(60115, "Password length is at least 8 and up to 24 characters")
	ErrChangeEmail                 = newErr(60116, "Email can't be changed")
	ErrInvalidMobile               = newErr(60117, "Please enter the correct mobile number")
	ErrInvalidEmail                = newErr(60118, "Please enter the correct email address")
	ErrKycStatus                   = newErr(60119, "Current kyc status, update info is not allowed")
	ErrIncompleteSecurityStrategy  = newErr(60120, "Incomplete verification item")
	ErrNoGoogleSecret              = newErr(60121, "Please get google secret first")
	ErrNULLMobileAreaCode          = newErr(60122, "Mobile area code is not null")
	ErrNeedMobileVerify            = newErr(60123, "User need mobile verify")
	ErrNeedEmailVerify             = newErr(60124, "User need email verify")
	ErrNeedFundPwdVerify           = newErr(60125, "User need fund password verify")
	ErrAlreadySetFundPassword      = NewErr(60126, "Fund password already set")
	ErrFundPasswordLen             = newErr(60127, "Fund password length must be 6 numbers")
	ErrDuplicateKycInfo            = newErr(60128, "Duplicate kyc info")
	ErrFundPasswordNumber          = newErr(60129, "Fund password must be number")
	ErrDuplicateEmail              = newErr(60130, "Duplicate email")
	ErrDuplicateMobile             = newErr(60131, "Duplicate mobile")
	ErrWithdrawStatus              = newErr(60132, "Your current status, withdraw is not allowed")
	ErrDuplicateCredentials        = newErr(60133, "Duplicate credentials")
	ErrInvalidGoogleSecret         = newErr(60134, "Invalid google secret")
	ErrNotBoundGoogle              = newErr(60135, "Google secret has't bound")
	ErrNeedSecurityVerify          = newErr(60136, "Need security verify")
	ErrInvalidWithdrawAmount       = newErr(60137, "Invalid withdraw amount")
	ErrBehaviorVerifyFail          = newErr(60138, "Behavior verify fail")
	ErrAccountFundTransfer         = newErr(60139, "account fund transfer error")
	ErrTransferAccountBalances     = newErr(60140, "transfer account balances not enough")
	ErrInventory                   = newErr(60141, "inventory not enough")
	ErrUnsupportTransferType       = newErr(60142, "transfer type is not support")
	ErrUnsupportTransferCurrency   = newErr(60143, "transfer currency is not support")
	ErrSubscriptionNotInProgress   = newErr(60144, "did not arrive at the subscription time")
	ErrSubscriptionAccountBalances = newErr(60145, "subscription account balances not enough")
	ErrInvalidDecimals             = newErr(60146, "invalid decimals")
	ErrInvalidTransferAccount      = newErr(60147, "Invalid transfer account")

	//钱包服务的错误码(60150-60200)
	ErrNoCurrency          = newErr(60150, "no the currency info")
	ErrNoWallet            = newErr(60151, "no the type wallet")
	ErrNoAddress           = newErr(60152, "no the currency address info")
	ErrInvalidAddress      = newErr(60153, "invalid address")
	ErrDisableWithdraw     = newErr(60154, "the currency has closed withdraw")
	ErrLTMinWithdrawAmount = newErr(60155, "the amount less than the min withdraw amount")
	ErrInvalidTag          = newErr(60156, "invalid tag info")
	ErrBlackUser           = newErr(60157, "black user")
	ErrInsufficientBalance = newErr(60158, "Insufficient balance")
	ErrEngineFreeze        = newErr(60159, "Engine freeze error")
	ErrEngineDeduction     = newErr(60160, "Engine deduction error")
	ErrEngineDeposit       = newErr(60161, "Engine deposit failed")
	ErrNoKycWithdraw       = newErr(60162, "kyc status no allowed to withdraw")

	//otc服务的错误码
	ErrUserKycNoAuth          = newErr(60163, "kyc is not certified")
	ErrUserNoPayment          = newErr(60164, "payment method not added")
	ErrCancelOrderLimit       = newErr(60165, "too many cancellations today")
	ErrExistUnFinishedOrder   = newErr(60166, "exist unfinished order")
	ErrAdvertisingWasModified = newErr(60167, "the AD was modified")
	ErrNotInHandleTime        = newErr(60168, "not currently in the merchant processing time range")
	ErrIncorrectOrderStatus   = newErr(60169, "incorrect order status")
	ErrIllegalOrderHanle      = newErr(60170, "illegal order handle")
	ErrSellerNoPay            = newErr(60171, "the seller has no supported collection method")
	ErrBuyerNoPay             = newErr(60172, "the buyer has no supported payment method")
	ErrTotalPriceOutOfRange   = newErr(60173, "the total price is out of range")
	ErrInsufficientInventory  = newErr(60174, "no more inventory")
	ErrInvalidOrder           = newErr(60175, "invalid order")
	ErrNoAllowCreateOrder     = newErr(60176, "current is not allow create order")
	ErrOtcAccountBalances     = newErr(60177, "otc account balance is enough")
	ErrInvalidPayMethod       = newErr(60178, "invalid pay method")
	ErrUnRelableAmount        = newErr(60179, "withdraw has unrelable amount")
	ErrInvalidPrecision       = newErr(60180, "invalid precision")
	ErrUnSupportCurrency      = newErr(60181, "unsupport currency")

	//服务之间调用错误码(60201-60300)
	ErrCaptcha = newErr(60201, "captcha service transfer error")
	ErrUser    = newErr(60202, "user service transfer error")
	ErrEngine  = newErr(60203, "engine service transfer error")

	//系统层面错误码(60301-63400)
	ErrSystemStopServing = newErr(60301, "System stop Serving")

	//行情错误码(60400-60500)
	ErrSymbolCollected         = newErr(60400, "symbol has been collected")
	ErrSymbolCollectedCanceled = newErr(60401, "symbol collected has been canceled")
	ErrSymbolNotCollected      = newErr(60402, "symbol not collected, can not canceled")

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

