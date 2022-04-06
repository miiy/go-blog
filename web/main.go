package main

import (
	"flag"
	"goblog.com/web/router"
)

func main() {
	//addr := flag.String("host", "127.0.0.1:8080", "host")
	conf := flag.String("c", "./config/default.yaml", "config file")
	flag.Parse()

	app, cleanUp, err := InitApplication(*conf)
	if err != nil {
		panic(err)
	}
	defer cleanUp()


	app.RegisterRouter(router.Router)

	// health
	// router: /health/*
	//InitializeHealth(app).RegisterRouter()
	app.Run(app.Config.App.Addr)
}