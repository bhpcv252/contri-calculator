build:
	@go build -o contri ./cmd/contri

run: build
	@./contri

test:
	@go test -v ./...
