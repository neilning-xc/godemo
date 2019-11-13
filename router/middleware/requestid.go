package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")

		if requestId == "" {
			v4, _ := uuid.NewV4()
			requestId = v4.String()
		}

		c.Set("X-Request-Id", requestId)
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
