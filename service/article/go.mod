module goblog.com/service/article

go 1.17

replace (
	goblog.com/api => ../../api
	goblog.com/pkg => ../../pkg
)

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	go.uber.org/zap v1.21.0
	goblog.com/api v0.0.0-00010101000000-000000000000
	goblog.com/pkg v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gorm.io/gorm v1.23.4
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/wire v0.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/net v0.0.0-20220407224826-aac1ed45d8e3 // indirect
	golang.org/x/sys v0.0.0-20220406163625-3f8b81556e12 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac // indirect
	gorm.io/driver/mysql v1.3.3 // indirect
)
