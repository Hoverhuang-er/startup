all:

debug:
	go run cmd/your_app_main/main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app cmd/your_app_main/main.go

gen_pb:
    protoc --go_out=./pkg/your_public_libs/pb --go_opt=paths=source_relative \
    --go-grpc_out=./pkg/your_public_libs/pb --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out ./pkg/your_public_libs/pb --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt standalone=true \
    --grpc-gateway_opt paths=source_relative \
    --openapiv2_out ./pkg/your_public_libs/pb  --openapiv2_opt logtostderr=true \
    ./startup.proto