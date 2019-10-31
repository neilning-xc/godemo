package main

import (
	"log"
	"net/http"

	"godemo/router"

	"github.com/julienschmidt/httprouter"
)

func NewRouters(routers router.Routers) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routers {
		handler := route.HandlerFunc
		router.Handle(route.Method, route.Path, handler)
	}

	return router
}

func main() {
	allRouters := router.AllRouters()
	router := NewRouters(allRouters)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
