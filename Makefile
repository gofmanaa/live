BINARY_NAME=go-live

all: dep build run

build:
#	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux main.go
#	GOARCH=amd64 GOOS=window go build -o ./bin/${BINARY_NAME}-windows main.go

run:
	./bin/${BINARY_NAME}-linux

dep:
	go mod tidy

test:
	go test ./...

clean:
	go clean
#	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
#	rm ./bin/${BINARY_NAME}-windows
