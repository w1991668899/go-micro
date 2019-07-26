package opentracing

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func GetClientFromContext(ctx context.Context, client *redis.Client) *redis.Client {
	if ctx == nil {
		return client
	}
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan == nil {
		return client
	}
	clientCopy := client.WithContext(ctx)
	clientCopy.WrapProcess(func(oldProcess func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
		return func(cmd redis.Cmder) error {
			tr := parentSpan.Tracer()
			sp := tr.StartSpan("redis", opentracing.ChildOf(parentSpan.Context()))
			defer sp.Finish()
			ext.DBType.Set(sp, "redis")
			sp.SetTag("db.method", cmd.Name())
			sp.SetTag("db.args", cmd.Args())
			return oldProcess(cmd)
		}
	})
	return clientCopy
}
