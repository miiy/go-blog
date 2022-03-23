protoc -I ./proto/v1 \
-I ../../third_party/googleapis \
-I ../../third_party/protoc-gen-validate \
--go_out ./proto/v1 --go_opt paths=source_relative \
--go-grpc_out ./proto/v1 --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./proto/v1 --grpc-gateway_opt paths=source_relative \
--validate_out="lang=go:./proto/v1" --validate_opt paths=source_relative \
./proto/v1/article_service.proto

protoc -I ./proto/v1 \
-I ../../third_party/googleapis \
--go_out ./proto/v1 --go_opt paths=source_relative \
--go-grpc_out ./proto/v1 --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./proto/v1 --grpc-gateway_opt paths=source_relative \
--openapiv2_out . \
--openapiv2_opt logtostderr=true \
--openapiv2_opt use_go_templates=true \
./proto/v1/article_service.proto