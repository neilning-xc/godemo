package main

import (
	"github.com/spf13/viper"
	"log"
	"net/http"

	"godemo/config"
	"godemo/model"
	"godemo/router"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "Config file path")
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
	pflag.Parse()

	allRouters := router.AllRouters()
	router := NewRouters(allRouters)

	if err := config.Init(*cfg); err != nil {
		log.Fatalf("Load config fail")
	}

	// 初始化数据库
	model.Init()

	err := http.ListenAndServe(viper.GetString("port"), router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
