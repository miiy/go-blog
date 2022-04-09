module goblog.com/service/feedback

go 1.17

replace (
	goblog.com/api => ../../api
	goblog.com/pkg => ../../pkg
)

require (
	goblog.com/api v0.0.0-00010101000000-000000000000
	goblog.com/pkg v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0 // indirect
	golang.org/x/net v0.0.0-20220407224826-aac1ed45d8e3 // indirect
	golang.org/x/sys v0.0.0-20220406163625-3f8b81556e12 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac // indirect
	google.golang.org/grpc v1.45.0 // indirect
)
