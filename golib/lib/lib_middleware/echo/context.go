package echomw

import (
	"context"
	"github.com/labstack/echo"
	"go-micro/golib/lib/lib_middleware/opentracing"
)

type EchoContextWrapper struct {
	echo.Context
}

func (c *EchoContextWrapper) ContextFromEcho() context.Context {
	span := c.Get(opentracing.SpanFromEchoKey)
	if span == nil {
		return context.Background()
	}
	return opentracing.ContextWithSpan(span)
}

func ContextWrapper() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &EchoContextWrapper{c}
			return h(cc)
		}
	}
}
