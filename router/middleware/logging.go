package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/willf/pad"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now().UTC()

		ip := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		c.Next()

		endTime := time.Now().UTC()
		latency := endTime.Sub(startTime)

		log.Infof("%-13s | %-12s | %s %s", latency, ip, pad.Right(method, 5, ""), path)
	}
}
