package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

func RegisterMetrics() error {
	if err := prometheus.Register(RequestsProcessingTime); err != nil {
		return err
	}
	return nil
}

var RequestsProcessingTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "request_processing_time_ms",
	Help:    "Histogram of response time for handler in mseconds",
	Buckets: []float64{1, 5, 10, 50, 100, 200, 500, 1000, 5000, 10000, 30000, 60000, 120000, 180000, 240000, 300000},
}, []string{"request"})

func SetRequestTime(request string, durationMS float64) {
	RequestsProcessingTime.With(prometheus.Labels{"request": request}).Observe(durationMS)
}
