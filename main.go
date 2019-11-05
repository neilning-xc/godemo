package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"godemo/config"
	"godemo/model"
	"godemo/router"

	"github.com/lexkong/log"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "Config file path")
)

func main() {
	// Read command line parmeter
	pflag.Parse()

	// Load configuration file
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	config.InitLog()

	g := gin.Default()

	// Load router
	router.Load(g)

	// Init database
	model.Init()
	defer model.CloseDB()

	// start http serser
	log.Infof("Starting listening requests on http server: %s", viper.GetString("url"))
	log.Info(http.ListenAndServe(viper.GetString("port"), g).Error())
}
