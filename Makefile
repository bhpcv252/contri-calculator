build:
	@go build -o contri . 

run: build
	@./contri

test:
	@go test -v ./...

