package main

import (
	"flag"
	"github.com/miiy/go-blog/pkg/api"
	"github.com/miiy/go-blog/pkg/config"
	"github.com/miiy/go-blog/pkg/database"
	"github.com/miiy/go-blog/pkg/migrate"
	"log"
	"os"
)

func main() {
	cFile := flag.String("-f", "./configs/default.yaml", "config path")
	flag.Parse()

	// config
	c, err := config.NewConfig(*cFile)
	if err != nil {
		log.Fatalln(err)
	}

	// db
	db, err := database.NewDatabase(&c.Database)
	if err != nil {
		log.Fatalln(err)
	}

	api.MigrationFiles()

	os.Exit(1)
	if err := migrate.Run(db.DB, "file:///migrations"); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}