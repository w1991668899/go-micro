package prometheus

import (
	"context"
	"github.com/labstack/echo"
	"github.com/micro/go-micro/server"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-micro/golib/lib/lib_config"
	"os"
	"strconv"
	"time"
)

var defaultLabelNames = []string{"node", "host", "status", "method", "handler"}

func APIMetricsByConfig(config lib_config.ConfPrometheus) echo.MiddlewareFunc {
	if config.Disable {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return next
		}
	}
	opts := make([]Option, 0)
	opts = append(opts, Namespace(config.Namespace))
	opts = append(opts, Subsystem(config.Subsystem))
	opts = append(opts, MetricsPath(config.MetricsPath))
	return APIMetrics(opts...)
}

func APIMetrics(options ...Option) echo.MiddlewareFunc {
	opts := initOptions(options...)
	// Counter
	requestCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Name:      "http_request_total",
			Help:      "Total request count.",
		},
		defaultLabelNames,
	)

	requestErrorCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Name:      "http_request_error_total",
			Help:      "Total request count.",
		},
		defaultLabelNames,
	)

	// Summary
	requestLatency := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Name:      "http_request_latency",
			Help:      "Request duration in milliseconds.",
		},
		defaultLabelNames,
	)

	responseSize := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: opts.Namespace,
			Subsystem: opts.Subsystem,
			Name:      "http_response_size",
			Help:      "Response size in bytes.",
		},
		defaultLabelNames,
	)

	opts.Registry.MustRegister(requestCount, requestErrorCount, requestLatency, responseSize)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "-"
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			// 拦截metrics path，默认"/metrics"
			if req.URL.Path == opts.MetricsPath {
				promhttp.Handler().ServeHTTP(c.Response(), c.Request())
				return nil
			}

			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}

			latency := time.Since(start)
			status := strconv.Itoa(res.Status)

			requestCount.WithLabelValues(hostname, req.Host, status, req.Method, req.RequestURI).Inc()
			requestErrorCount.WithLabelValues(hostname, req.Host, status, req.Method, req.RequestURI).Inc()
			requestLatency.WithLabelValues(hostname, req.Host, status, req.Method, req.RequestURI).Observe(float64(latency.Nanoseconds() / 1e6))
			responseSize.WithLabelValues(hostname, req.Host, status, req.Method, req.RequestURI).Observe(float64(res.Size))

			return nil
		}
	}
}

func MicroAPIMetrics() server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			return nil
		}
	}
}
