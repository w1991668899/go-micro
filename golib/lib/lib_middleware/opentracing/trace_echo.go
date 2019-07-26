package opentracing

import (
	"context"

	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
)

const (
	SpanFromEchoKey = "echo_span"
	TraceId         = "trace_id"
)

func OpenTracing() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var span opentracing.Span
			opName := c.Request().Method + " " + c.Request().URL.Path

			wireContext, err := opentracing.GlobalTracer().Extract(
				opentracing.HTTPHeaders,
				opentracing.HTTPHeadersCarrier(c.Request().Header))
			if err != nil {
				// 启动新Span
				span = opentracing.StartSpan(opName)
			} else {
				span = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
			}

			defer span.Finish()
			c.Set(SpanFromEchoKey, span)
			c.Set(TraceId, TraceIdFromSpan(span))
			span.SetTag("component", "echo api")
			span.SetTag("span.kind", "server")
			span.SetTag("http.url", c.Request().Host+c.Request().RequestURI)
			span.SetTag("http.method", c.Request().Method)

			if err := next(c); err != nil {
				span.SetTag("error", true)
				c.Error(err)
			}

			span.SetTag("error", false)
			span.SetTag("http.status_code", c.Response().Status)

			return nil
		}
	}
}

func ContextFromEcho(ctx echo.Context) context.Context {
	span := ctx.Get(SpanFromEchoKey)
	if span == nil {
		return context.Background()
	}
	return opentracing.ContextWithSpan(context.Background(), span.(opentracing.Span))
}

func ContextWithSpan(span interface{}) context.Context {
	s, ok := span.(opentracing.Span)
	if ok {
		return opentracing.ContextWithSpan(context.Background(), s)
	}
	return context.Background()
}
