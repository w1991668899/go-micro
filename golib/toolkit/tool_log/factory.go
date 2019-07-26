package tool_log

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/common/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger *zap.Logger
)

func init() {
	var err error
	zapLogger, err = zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	if err != nil{
		log.Fatalln("zap log init fail: ", err)
	}
}

func LogWith(fields ...zap.Field) *zap.Logger {
	return zapLogger.With(fields...)
}

type Factory struct {
	logger *zap.Logger
}

func NewFactory(logger *zap.Logger) Factory {
	return Factory{logger: logger}
}

// background 返回一个不包含ctx的logger
func (f Factory) Background() Logger {
	return logger{logger: f.logger}
}

// 如果context不包含openTracing的span, 返回background
func (f Factory) For(ctx context.Context) Logger {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		// TODO for Jaeger span extract trace/span IDs as fields
		return spanLogger{span: span, logger: f.logger}
	}
	return f.Background()
}

func (f Factory) With(fields ...zapcore.Field) Factory {
	return Factory{logger: f.logger.With(fields...)}
}
