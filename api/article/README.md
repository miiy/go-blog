protoc -I ./article/v1 \
-I ../third_party/googleapis \
--go_out ./article/v1 --go_opt paths=source_relative \
--go-grpc_out ./article/v1 --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./article/v1 --grpc-gateway_opt paths=source_relative \
--openapiv2_out ./article \
--openapiv2_opt logtostderr=true \
--openapiv2_opt use_go_templates=true \
./article/v1/article_service.proto