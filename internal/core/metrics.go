package core

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HttpMetrics struct {
	httpRequestsTotal *prometheus.CounterVec
	httpResponseTime  *prometheus.HistogramVec
}

type HttpMetricsData struct {
	m     *HttpMetrics
	timer *prometheus.Timer
}

func NewHttpMetrics() *HttpMetrics {
	return &HttpMetrics{
		httpRequestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "path", "status"},
		),
		httpResponseTime: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "http_response_time_seconds",
				Help: "HTTP response time in seconds",
			},
			[]string{"method", "path", "status"},
		),
	}
}

func (m *HttpMetrics) StartRequestMetrics(statusCode int, method string, path string) HttpMetricsData {
	status := strconv.Itoa(statusCode)

	labels := prometheus.Labels{
		"method": method,
		"path":   path,
		"status": status,
	}

	m.httpRequestsTotal.With(labels).Inc()
	timer := prometheus.NewTimer(m.httpResponseTime.With(labels))

	return HttpMetricsData{
		m:     m,
		timer: timer,
	}
}

func (m *HttpMetricsData) EndRequestMetrics() {
	m.timer.ObserveDuration()
}

func RegisterMetricsAt(app fiber.Router, url string, handlers ...fiber.Handler) {
	h := append(handlers, adaptor.HTTPHandler(promhttp.Handler()))
	app.Get(url, h...)
}
