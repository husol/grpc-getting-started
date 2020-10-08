.PHONY: build run client gen-protoc

run:
	go run cmd/main.go

client:
	go run cmd/client.go

gen-protoc:
	@protoc --go_out=plugins=grpc:./proto --proto_path=proto --go_opt=paths=source_relative proto/service.proto
