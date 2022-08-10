all: clean build

build:
	go mod tidy
	go build -o build/main main.go

clean:
	go clean
	rm -rf build
