build:
	@go build -o bin/contri-calculator

run:
	@go run .

test:
	@go test -v ./...

