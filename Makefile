build:
	@echo "building binary..."
	@go build -o bin/goshort cmd/server/main.go

run-bin:
	@echo "running binary..."
	@./bin/goshort

run-dev:
	@echo "running dev..."
	@go run cmd/server/main.go

