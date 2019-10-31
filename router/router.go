package router

import (
	"godemo/controllers/user"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routers []Route

func AllRouters() Routers {
	routers := Routers{
		Route{
			Name:        "user.list",
			Method:      "GET",
			Path:        "/user",
			HandlerFunc: user.List,
		},
	}

	return routers
}
