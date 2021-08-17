package api

import (
	"embed"
	"fmt"
	"io/fs"

	//"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed */migrations/*.sql
var mfs embed.FS

func MigrationFiles() {
	d, _ := fs.ReadFile(mfs, ".")
	fmt.Println(d)
	for _, v := range d {
		fmt.Println(v)
	}
	//iofs.
	//iofs.New()
}