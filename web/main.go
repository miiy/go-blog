package main

import (
	"flag"
	"goblog.com/web/router"
)

func main() {
	//host := flag.String("host", "127.0.0.1", "host")
	conf := flag.String("c", "./config/default.yaml", "config file")
	flag.Parse()

	app, f, err := InitApplication(*conf)
	if err != nil {
		panic(err)
	}
	defer f()

	app.RegisterRouter(router.Router)
	app.Router.Run(app.Config.App.Addr)
}