LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=$(CURDIR)/migrations
LOCAL_MIGRATION_DSN="host=localhost port=54321 dbname=chat-server-db user=postgres password=postgres sslmode=disable"

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-chat-api

generate-chat-api:
	mkdir -p pkg/chat_v1
	protoc --proto_path api/chat_v1 \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/chat_v1/chat.proto

build:
	GOOS=linux GOARCH=amd64 go build -o chat-server cmd/main.go

copy-to-server:
	scp chat-server v_voronin@95.174.94.156:

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t pgregistry.cr.cloud.ru/chat-server:v0.0.1 .
	docker login pgregistry.cr.cloud.ru -u f921987766c413d6b9e152f97ef15b5d -p eddce5e8de0798cd6ce2319f2cd2b8d9
	docker push pgregistry.cr.cloud.ru/chat-server:v0.0.1

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v