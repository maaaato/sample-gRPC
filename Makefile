protoc-echo:
	protoc --proto_path ./proto \
	--go_out=plugins=grpc:./proto \
	--go_opt=module=github.com/maaaato/sample-gRPC/proto \
	./proto/echo.proto