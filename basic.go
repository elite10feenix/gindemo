package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

func main() {
	router := gin.Default()

	router.Use(LoggerMiddlware())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	router.Run(":8080")
}
