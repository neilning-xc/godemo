package router

import (
	"godemo/controllers/user"
	"godemo/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Url does't exist")
	})

	g.NoMethod(func(c *gin.Context) {
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	})

	g.POST("/login", user.Login)

	v1 := g.Group("/v1")
	v1.Use(middleware.Auth())
	{
		v1.GET("/user", user.List)
		v1.POST("/user", user.Create)
	}

	return g
}
