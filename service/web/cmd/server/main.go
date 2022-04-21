package main

import (
	"flag"
	"fmt"
	"goblog.com/service/web/internal/router"
)

func main() {
	port := flag.Int("port", 8080, "host")
	conf := flag.String("c", "./configs/default.yaml", "config file")
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
	app.Run(fmt.Sprintf("0.0.0.0:%d", *port))
}