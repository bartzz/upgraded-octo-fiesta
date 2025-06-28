.PHONY: start-api stop-api fmt fmt-go fmt-golangci

fmt: fmt-go fmt-goimports fmt-golangci

start-api:
	@echo "Starting API using docker compose"
	@docker compose up -d
	@docker compose logs -f

stop-api:
	@echo "Stopping API"
	@docker compose down

fmt-go:
	@echo "→ go fmt"
	@go fmt ./...

fmt-goimports:
	@echo "→ goimports"
	@goimports -w .

fmt-golangci:
	@echo "→ golangci-lint fixes"
	@golangci-lint run --fix ./...