TARGET=server

proto:
#protoc --proto_path=api/tag/v1 --proto_path=third_party/proto --go_out=plugins=grpc:api/tag/v1 --go_opt paths=source_relative tag_service.proto

# example proto
protoc -I ./pkg/api/example/proto/v1 -I ./third_party/proto \
       -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.1 \
       --go_out ./pkg/api/example/proto/v1 --go_opt paths=source_relative \
       --go-grpc_out ./pkg/api/example/proto/v1 --go-grpc_opt paths=source_relative \
       --grpc-gateway_out ./pkg/api/example/proto/v1 --grpc-gateway_opt paths=source_relative \
       --validate_out="lang=go:./pkg/api/example/proto/v1" \
       ./pkg/api/example/proto/v1/example_service.proto

# auth proto
protoc -I ./pkg/api/auth/proto -I ./third_party/proto \
       --go_out ./pkg/api/auth/proto --go_opt paths=source_relative \
       --go-grpc_out ./pkg/api/auth/proto --go-grpc_opt paths=source_relative \
       --grpc-gateway_out ./pkg/api/auth/proto --grpc-gateway_opt paths=source_relative \
       ./pkg/api/auth/proto/auth_service.proto

# tag proto
protoc -I ./pkg/api/tag/proto/v1 -I ./third_party/proto \
       --go_out ./pkg/api/tag/proto/v1 --go_opt paths=source_relative \
       --go-grpc_out ./pkg/api/tag/proto/v1 --go-grpc_opt paths=source_relative \
       --grpc-gateway_out ./pkg/api/tag/proto/v1 --grpc-gateway_opt paths=source_relative \
       ./pkg/api/tag/proto/v1/tag_service.proto

# userpost proto
protoc -I ./pkg/api/userpost/proto \
       -I ./third_party/proto \
       -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.1 \
       --go_out ./pkg/api/userpost/proto --go_opt paths=source_relative \
       --go-grpc_out ./pkg/api/userpost/proto --go-grpc_opt paths=source_relative \
       --grpc-gateway_out ./pkg/api/userpost/proto --grpc-gateway_opt paths=source_relative \
       --validate_out="lang=go:./pkg/api/userpost/proto" --validate_opt  paths=source_relative \
       ./pkg/api/userpost/proto/userpost.proto

# usertag proto
protoc -I ./pkg/api/usertag/proto \
       -I ./third_party/proto \
       -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.1 \
       --go_out ./pkg/api/usertag/proto --go_opt paths=source_relative \
       --go-grpc_out ./pkg/api/usertag/proto --go-grpc_opt paths=source_relative \
       --grpc-gateway_out ./pkg/api/usertag/proto --grpc-gateway_opt paths=source_relative \
       --validate_out="lang=go:./pkg/api/usertag/proto" --validate_opt  paths=source_relative \
       ./pkg/api/usertag/proto/usertag.proto

# feedback proto
protoc -I ./pkg/api/feedback/proto \
       -I ./third_party/proto \
       -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.1 \
       --go_out ./pkg/api/feedback/proto --go_opt paths=source_relative \
       --go-grpc_out ./pkg/api/feedback/proto --go-grpc_opt paths=source_relative \
       --grpc-gateway_out ./pkg/api/feedback/proto --grpc-gateway_opt paths=source_relative \
       --validate_out="lang=go:./pkg/api/feedback/proto" --validate_opt  paths=source_relative \
       ./pkg/api/feedback/proto/feedback.proto


BINARY="go-web"

.PHONY: default
default:
	@go build -o ${BINARY} ./cmd/server

.PHONY: clean
clean:
	@rm -rf ${BINARY}

.PHONY: help
help:
	@echo "make:       build app"
	@echo "make clean: clean binary file"
