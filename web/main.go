package main

import (
	"flag"
	"fmt"
	"goblog.com/pkg/viper"
	"goblog.com/web/pkg/config"
	"goblog.com/web/router"
	"log"
)

func main() {
	configFile := flag.String("c", "./config/default.yaml", "config file")
	flag.Parse()

	v, err := viper.NewViper("default.yaml", "yaml", []string{"./config"})
	if err != nil {
		panic(err)
	}
	fmt.Println(v.AllKeys())
	fmt.Println(v.Get("app.name"))

	// config
	c, err := config.NewConfig(*configFile)
	if err != nil {
		log.Fatalln(err)
	}

	app, f, err := InitApplication()
	if err != nil {
		panic(err)
	}
	defer f()

	router.RegisterRouter(app.Router, v)

	app.Router.Run(c.App.Addr)
}