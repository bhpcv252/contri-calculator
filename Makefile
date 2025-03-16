VERSION = $(shell git describe --tags || echo "dev")

build:
	@go build -ldflags "-X main.version=$(VERSION)" -o contri .

run: build
	@./contri

test:
	@go test -v ./...
