package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"goblog.com/web/pkg/config"
	"goblog.com/web/router"
	"log"
)

func main() {
	configFile := flag.String("f", "config/default.yaml", "config file")
	flag.Parse()

	// config
	c, err := config.NewConfig(*configFile)
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	router.RegisterRouter(r, c.App)
	if c.App.Debug {
		gin.SetMode(gin.DebugMode)
	}

	r.Run(c.App.Addr)
}