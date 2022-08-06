BINARY_NAME=railcar

build:
	GOARCH=amd64 GOOS=darwin go build -ldflags="-X 'pkg/build.sha1=$(shell git rev-parse HEAD)' -X 'railcar/pkg/build.buildTime=$(shell date)'" -o ./bin/${BINARY_NAME}-darwin cmd/server/main.go
	# GOARCH=amd64 GOOS=linux go build -ldflags='-X pkg/build.sha1=$(shell git rev-parse HEAD) -X "railcar/pkg/build.buildTime=$(shell date)"' -o ./bin/${BINARY_NAME}-linux cmd/server/main.go
	go test ./... 

run:
	./bin/${BINARY_NAME}-darwin

build_and_run: build run

test: 
clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
	# rm ./bin/${BINARY_NAME}-windows
