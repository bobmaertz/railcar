BINARY_NAME=railcar

all: test build run 

build: 
	GOARCH=amd64 GOOS=darwin go build \
		-ldflags="-X 'github.com/bobmaertz/railcar/pkg/build.sha1=$(shell git rev-parse HEAD)' -X 'github.com/bobmaertz/railcar/pkg/build.buildTime=$(shell date)'" \
		-o ./bin/${BINARY_NAME}-darwin cmd/server/main.go

run:
	./bin/${BINARY_NAME}-darwin

test: 
	go test ./... -coverprofile=coverage.out


build_and_run: build run

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	# rm ./bin/${BINARY_NAME}-linux
	# rm ./bin/${BINARY_NAME}-windows
