package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type LoggerMiddleware struct{}

func (l *LoggerMiddleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		t := time.Now()

		c.Next()

		// after request
		latency := time.Since(t)
		log.Printf("Letency: %s", latency)
	}
}
