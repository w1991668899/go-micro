package opentracing

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/registry"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

// accept an opentracing Tracer and returns a micro client Call Wrapper
func NewCallWrapper(ot opentracing.Tracer) client.CallWrapper {
	return func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
			name := fmt.Sprintf("%s.%s", req.Service(), req.Method())

			var parentCtx opentracing.SpanContext
			if parent := opentracing.SpanFromContext(ctx); parent != nil {
				parentCtx = parent.Context()
			}
			span := ot.StartSpan(
				name,
				opentracing.ChildOf(parentCtx),
				ext.SpanKindRPCClient,
			)

			defer span.Finish()
			md, ok := metadata.FromContext(ctx)
			if !ok {
				md = make(map[string]string)
			}

			if err := opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md)); err != nil {
				return err
			}
			ctx = opentracing.ContextWithSpan(ctx, span)
			ctx = metadata.NewContext(ctx, md)
			err := cf(ctx, node, req, rsp, opts)
			if err != nil {
				span.SetTag("error", true)
				span.LogFields(
					log.String("event", "call error"),
					log.Object("error", err),
				)
			}
			return err
		}
	}
}

// accept an opentracing Tracer and returns a micro server Handler Wrapper
func NewHandlerWrapper(ot opentracing.Tracer) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			name := fmt.Sprintf("%s.%s", req.Service(), req.Method())
			ctx, span, err := traceIntoContext(ctx, ot, name)
			if err != nil {
				return err
			}
			defer span.Finish()
			return h(ctx, req, rsp)
		}
	}
}

func traceIntoContext(ctx context.Context, tracer opentracing.Tracer, name string) (context.Context, opentracing.Span, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, err := tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	if err != nil {
		sp = tracer.StartSpan(name)
	} else {
		sp = tracer.StartSpan(name, opentracing.ChildOf(wireContext))
	}

	ext.SpanKindRPCServer.Set(sp)

	ctx = opentracing.ContextWithSpan(ctx, sp)
	ctx = metadata.NewContext(ctx, md)
	return ctx, sp, nil
}
