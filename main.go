package main

import (
	"log"
	"net/http"

	"godemo/model"
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

	// 初始化数据库
	model.Init()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
