package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
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

func proxy(c *gin.Context) {
	remote, err := url.Parse("http://localhost:8082")
	if err != nil {
		c.String(http.StatusInternalServerError, "URL Parsing error")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		proxyPath := c.Param("proxyPath")
		if proxyPath == "" {
			proxyPath = "/"
		}
		req.URL.Path = strings.TrimSuffix(remote.Path, "/") + proxyPath
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	r := gin.Default()
	r.Use(RateLimiter())
	r.Any("/proxy/*proxyPath", proxy)
	r.Run(":8081")
}
