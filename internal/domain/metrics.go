package domain

type HttpMetrics interface {
	StartRequestMetrics(statusCode int, method string, path string)
	EndRequestMetrics()
}
