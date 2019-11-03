package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"

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
	// Read command line parmeter
	pflag.Parse()

	// Load router
	allRouters := router.AllRouters()
	router := NewRouters(allRouters)

	// Load configuration file
	if err := config.Init(*cfg); err != nil {
		log.Fatalf("Load config fail")
	}

	// Init database
	model.Init()

	// start http serser
	err := http.ListenAndServe(viper.GetString("port"), router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
