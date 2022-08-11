all: clean build

build:
	go mod tidy
	go build -o build/main main.go

clean:
	go clean
	rm -rf build

gen:
	protoc -I=. --go_opt=paths=source_relative --go_out=. \
		--go-grpc_opt=paths=source_relative --go-grpc_out=. \
		protocol/protocol.proto
