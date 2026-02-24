.PHONY: proto build tidy

proto:
	@mkdir -p api/gen/seat
	protoc --proto_path=api/proto \
		--go_out=api/gen/seat --go_opt=paths=source_relative \
		--go-grpc_out=api/gen/seat --go-grpc_opt=paths=source_relative \
		api/proto/seat.proto

tidy:
	go mod tidy

build:
	go build -o bin/extractor cmd/extractor/main.go
