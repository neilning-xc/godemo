package middleware

import (
	"godemo/controllers"
	"godemo/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			controllers.SendJSONResponse(c, http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
