.PHONY: run
run: 
	@go run cmd/main.go

.PHONY: build
build:
	@go build -o bin/main cmd/main.go