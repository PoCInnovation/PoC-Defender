package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"net/http/httputil"
	"strings"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"

	"MyRevProxy/config"
)

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(1, 4)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Limite exceed",
			})
		}
	}
}
  
func proxy(c *gin.Context, config *config.Config) {
	proxy := httputil.NewSingleHostReverseProxy(config.BackendURL)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		proxyPath := c.Param("proxyPath")
		if proxyPath == "" {
			proxyPath = "/"
		}
		req.URL.Path = strings.TrimSuffix(config.BackendURL.Path, "/") + proxyPath
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}
	initMetrics()
	r := gin.Default()
	r.Use(RateLimiter())
	r.Use(MetricsMiddleware())
	RegisterMetricsEndpoint(r)
	r.Any("/proxy/*proxyPath", func(c *gin.Context) {
		proxy(c, config)
	})
	port := strconv.Itoa(config.ServerPort)
	r.Run(":" + port)
}
