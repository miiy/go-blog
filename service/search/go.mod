module goblog.com/service/search

go 1.17

replace (
	goblog.com/api => ../../api
	goblog.com/pkg => ../../pkg
)

require (
	github.com/elastic/go-elasticsearch/v8 v8.1.0
	goblog.com/pkg v0.0.0-00010101000000-000000000000
)

require github.com/elastic/elastic-transport-go/v8 v8.1.0 // indirect
