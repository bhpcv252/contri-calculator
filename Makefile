build:
	@go build -o bin/contri-calculator

run: build
	@./bin/contri-calculator

test:
	@go test -v ./...

