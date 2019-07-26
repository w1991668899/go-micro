package opentracing

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	"go-micro/golib/toolkit/tool_log"
	"go-micro/golib/toolkit/tool_metrics"
	"io"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"time"
)

var (
	Logger tool_log.Factory
)

func NewTracerByConfig(config ConfigJaeger) (opentracing.Tracer, io.Closer) {
	return NewTracer(
		ServiceName(config.ServiceName),
		AgentAddr(config.AgentAddr),
		Disable(config.Disable),
	)
}

// Init creates a new instance of Jaeger tracer.
func NewTracer(opts ...Option) (opentracing.Tracer, io.Closer) {

	option := Options{}
	for _, opt := range opts {
		opt(&option)
	}

	//init tool_log
	zapLogger := tool_log.LogWith(zap.String("service", option.ServiceName))
	Logger = tool_log.NewFactory(zapLogger)

	// init metics
	metricsFactory := tool_metrics.GetMetrics()
	metricsFactory.Namespace(option.ServiceName, nil)

	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		Logger.Background().Fatal("cannot parse Jaeger env vars", zap.Error(err))
	}
	cfg.ServiceName = option.ServiceName
	cfg.Disabled = option.Disable
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter.LocalAgentHostPort = option.AgentAddr
	//cfg.Reporter.LogSpans = true

	//a quick hack to ensure random generators get different seeds, which are based on current time.
	time.Sleep(100 * time.Millisecond)
	jaegerLogger := jaegerLoggerAdapter{Logger.Background()}

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jaegerLogger),
		jaegercfg.Metrics(metricsFactory),
		jaegercfg.Observer(rpcmetrics.NewObserver(metricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)
	if err != nil {
		Logger.Background().Fatal("cannot initialize Jaeger Tracer", zap.Error(err))
	}

	opentracing.SetGlobalTracer(tracer)

	return tracer, closer
}

// call this function must Init first
func GetTracer() opentracing.Tracer {
	return opentracing.GlobalTracer()
}

func TraceIdFromContext(ctx context.Context) (traceId string) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		traceId = TraceIdFromSpan(span)
	}
	return
}

func TraceIdFromSpan(span opentracing.Span) string {
	s, ok := span.Context().(jaeger.SpanContext)
	if ok {
		return s.TraceID().String()
	}
	return ""
}

func StartSpan(ctx context.Context, operationName string, tags map[string]interface{}) (span opentracing.Span, traceId string) {
	var parentCtx opentracing.SpanContext
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		parentCtx = parent.Context()
	}

	span = opentracing.GlobalTracer().StartSpan(operationName, opentracing.ChildOf(parentCtx))
	for k, v := range tags {
		span.SetTag(k, v)
	}
	ext.SpanKindRPCClient.Set(span)
	traceId = TraceIdFromSpan(span)
	ctx = opentracing.ContextWithSpan(ctx, span)
	return span, traceId
}

func SpanIdFormSpan(span opentracing.Span) string {
	s, ok := span.Context().(jaeger.SpanContext)
	if ok {
		return s.SpanID().String()
	}
	return ""
}

type jaegerLoggerAdapter struct {
	logger tool_log.Logger
}

func (l jaegerLoggerAdapter) Error(msg string) {
	l.logger.Error(msg)
}

func (l jaegerLoggerAdapter) Infof(msg string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(msg, args...))
}
