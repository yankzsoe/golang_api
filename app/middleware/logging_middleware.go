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
		// only processes that last longer than 5 seconds will be written in the log.
		if latency.Seconds() > 5 {
			log.Printf("Letency: %v", latency)
		}
	}
}
