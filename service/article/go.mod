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
	goblog.com/pkg v0.0.0-00010101000000-000000000000
	google.golang.org/genproto v0.0.0-20220405205423-9d709892a2bf
	google.golang.org/grpc v1.45.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.28.0
	gorm.io/gorm v1.23.3
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/wire v0.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	goblog.com/api v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220307203707-22a9840ba4d7 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.3.2 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
