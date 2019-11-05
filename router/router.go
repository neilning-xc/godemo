package router

import (
	"godemo/controllers/user"
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

	v1 := g.Group("/v1")
	{
		v1.GET("/user", user.List)
	}

	return g
}
