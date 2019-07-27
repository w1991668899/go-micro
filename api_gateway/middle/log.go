package middleware

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"go-micro/golib/lib/lib_log"
	"io"
	"io/ioutil"
)

// 不要用
func Log() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			req := ctx.Request()

			var b bytes.Buffer
			_, err = io.Copy(&b, req.Body)
			if err != nil {
				lib_log.LogErrorLn(err)
			}
			rr, err := ioutil.ReadAll(&b)
			if err != nil {
				lib_log.LogErrorLn(err)
			} else {
				lib_log.LogInfo(logrus.Fields{
					"method": req.Method,
					"body":   string(rr),
				}, "request from client")
			}

			return next(ctx)
		}
	}
}
