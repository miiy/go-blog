package go_blog

import (
	"embed"
)

//go:embed migrations/*.sql
var MigrationFS embed.FS
