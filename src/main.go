package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"

	"MyRevProxy/config"
)

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
	r.Use(MetricsMiddleware())
	RegisterMetricsEndpoint(r)
	r.Any("/proxy/*proxyPath", func(c *gin.Context) {
		proxy(c, config)
	})
	port := strconv.Itoa(config.ServerPort)
	r.Run(":" + port)
}
