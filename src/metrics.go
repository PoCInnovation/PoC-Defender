package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "test_http_requests_total",
			Help: "Total number of HTTP requests",
		},
	)
)

func initMetrics() {
	prometheus.MustRegister(requestsTotal)
}

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestsTotal.Inc()
		c.Next()
	}
}

func RegisterMetricsEndpoint(r *gin.Engine) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
