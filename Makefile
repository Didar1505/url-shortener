LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-shortener-api

generate-shortener-api:
	mkdir -p pkg/url_shortener_v1
	protoc --proto_path ./proto --proto_path . \
	--go_out=pkg/url_shortener_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go.exe \
	--go-grpc_out=pkg/url_shortener_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc.exe \
	proto/urlshortener.proto

build-windows:
	go build -o ./bin/ ./url-service/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o service_linux ./url-service/main.go