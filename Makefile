all: clean gen build

build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/client cmd/client/main.go

clean:
	go clean
	rm -rf bin

gen:
	protoc -I=. --go_opt=paths=source_relative --go_out=. \
		--go-grpc_opt=paths=source_relative --go-grpc_out=. \
		protocol/protocol.proto
