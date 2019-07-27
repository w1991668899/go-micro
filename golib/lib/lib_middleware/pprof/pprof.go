package pprof

import (
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

func Serve() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
			if strings.HasPrefix(path, "/debug/pprof/") {
				http.DefaultServeMux.ServeHTTP(c.Response(), c.Request())
			} else {
				return next(c)
			}
			return nil
		}
	}
}
