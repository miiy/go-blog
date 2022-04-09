protoc -I ./article/v1 \
-I ../third_party/googleapis \
--go_out ./article/v1 --go_opt paths=source_relative \
--go-grpc_out ./article/v1 --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./article/v1 --grpc-gateway_opt paths=source_relative \
--openapiv2_out ./article \
--openapiv2_opt logtostderr=true \
--openapiv2_opt use_go_templates=true \
./article/v1/article.proto

protoc -I ./auth/v1 \
-I ../third_party/googleapis \
--go_out ./auth/v1 --go_opt paths=source_relative \
--go-grpc_out ./auth/v1 --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./auth/v1 --grpc-gateway_opt paths=source_relative \
--openapiv2_out ./auth \
--openapiv2_opt logtostderr=true \
--openapiv2_opt use_go_templates=true \
./auth/v1/auth.proto

protoc -I ./book/v1 \
-I ../third_party/googleapis \
--go_out ./book/v1 --go_opt paths=source_relative \
--go-grpc_out ./book/v1 --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./book/v1 --grpc-gateway_opt paths=source_relative \
--openapiv2_out ./book \
--openapiv2_opt logtostderr=true \
--openapiv2_opt use_go_templates=true \
./book/v1/book.proto

protoc -I ./feedback/v1 \
-I ../third_party/googleapis \
--go_out ./feedback/v1 --go_opt paths=source_relative \
--go-grpc_out ./feedback/v1 --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./feedback/v1 --grpc-gateway_opt paths=source_relative \
--openapiv2_out ./feedback \
--openapiv2_opt logtostderr=true \
--openapiv2_opt use_go_templates=true \
./feedback/v1/feedback.proto

protoc -I ./tag/v1 \
-I ../third_party/googleapis \
--go_out ./tag/v1 --go_opt paths=source_relative \
--go-grpc_out ./tag/v1 --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./tag/v1 --grpc-gateway_opt paths=source_relative \
--openapiv2_out ./tag \
--openapiv2_opt logtostderr=true \
--openapiv2_opt use_go_templates=true \
./tag/v1/tag.proto


protoc -I ./user/v1 \
-I ../third_party/googleapis \
--go_out ./user/v1 --go_opt paths=source_relative \
--go-grpc_out ./user/v1 --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./user/v1 --grpc-gateway_opt paths=source_relative \
--openapiv2_out ./user \
--openapiv2_opt logtostderr=true \
--openapiv2_opt use_go_templates=true \
./user/v1/user.proto


protoc -I ./userpost/v1 \
-I ../third_party/googleapis \
--go_out ./userpost/v1 --go_opt paths=source_relative \
--go-grpc_out ./userpost/v1 --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./userpost/v1 --grpc-gateway_opt paths=source_relative \
--openapiv2_out ./userpost \
--openapiv2_opt logtostderr=true \
--openapiv2_opt use_go_templates=true \
./userpost/v1/userpost.proto
