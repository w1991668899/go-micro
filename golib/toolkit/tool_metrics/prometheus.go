package tool_metrics

import "github.com/uber/jaeger-lib/metrics/prometheus"

var (
	metrics *prometheus.Factory
)

func init() {
	metrics = prometheus.New()
}

type MetricsFactory struct {
	metrics *prometheus.Factory
}

func GetMetrics() *prometheus.Factory {
	return metrics
}
