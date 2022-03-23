package openapi

import (
	"embed"
)

//go:embed *.json
var OpenAPIFS embed.FS