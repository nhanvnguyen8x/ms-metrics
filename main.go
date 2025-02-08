package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"ms-metrics/middleware"
	"net/http"
)

func main() {
	app := gin.Default()
	app.Use(middleware.PrometheusMiddleware())

	app.GET("/metrics", gin.WrapH(promhttp.Handler()))

	app.Use(func(c *gin.Context) {
		middleware.ActiveConnections.Inc()
		defer middleware.ActiveConnections.Dec()
		c.Next()
	})

	app.POST("/register", func(c *gin.Context) {
		middleware.UserRegistrations.Inc()
		c.JSON(http.StatusOK, gin.H{"message": "User registered"})
	})

	app.POST("/upload", func(c *gin.Context) {
		bodySize := c.Request.ContentLength
		middleware.RequestSizeHistogram.Observe(float64(bodySize))
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Uploaded %d bytes", bodySize)})
	})

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	if err := app.Run(":8080"); err != nil {
		log.Printf("Error in running server: %s", err.Error())
		return
	}
}
