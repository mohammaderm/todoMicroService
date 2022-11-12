package monitoring

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var defaultBuckets = []float64{.005, .02, .04, .07, .1, .15, .25, 0.5, 0.75, 1, 2, 3, 5, 10, 15, 20, 25, 30}

type MetricsCallectors interface {
	HttpResponseTime(method, path string, statusCode int, duration time.Duration)
	HttpRequestCount(method, path string, statusCode int)
	Start() error
}

type PrometheusMetrics struct {
	httpResponseTimeMetric *prometheus.HistogramVec
	httpRequestCountMetric *prometheus.CounterVec
	registry               *prometheus.Registry
	server                 *http.Server
}

func New(port string) MetricsCallectors {
	prometheusMetrics := &PrometheusMetrics{

		httpResponseTimeMetric: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "http_response_time_seconds",
			Help:    "response time of http endpoints in seconds",
			Buckets: defaultBuckets,
		}, []string{"method", "path", "status_code"}),

		httpRequestCountMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "http_request_count",
			Help: "Number of all http request",
		}, []string{"method", "path", "status_code"}),

		registry: prometheus.NewRegistry(),

		server: &http.Server{
			Addr:         ":" + port,
			ReadTimeout:  time.Second * 2,
			WriteTimeout: time.Second * 5,
		},
	}
	prometheusMetrics.registry.MustRegister(prometheusMetrics.httpResponseTimeMetric)
	prometheusMetrics.registry.MustRegister(prometheusMetrics.httpRequestCountMetric)

	return prometheusMetrics
}

func (p *PrometheusMetrics) HttpResponseTime(method, path string, statusCode int, duration time.Duration) {
	p.httpResponseTimeMetric.With(prometheus.Labels{
		"method": method, "path": path, "status_code": strconv.Itoa(statusCode),
	}).Observe(duration.Seconds())
}

func (p *PrometheusMetrics) HttpRequestCount(method, path string, statusCode int) {
	p.httpRequestCountMetric.With(prometheus.Labels{
		"method": method, "path": path, "status_code": strconv.Itoa(statusCode),
	}).Inc()
}

func (p PrometheusMetrics) Start() error {
	handler := promhttp.HandlerFor(p.registry, promhttp.HandlerOpts{})
	p.server.Handler = handler
	if err := p.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
