.PHONY: help infra-up infra-down dev test fmt-check lint tidy

help:
	@echo "Targets:"
	@echo "  make infra-up     Start postgres + redis"
	@echo "  make infra-down   Stop postgres + redis"
	@echo "  make dev          Run API locally (loads .env)"
	@echo "  make test         Run Go tests"
	@echo "  make fmt-check    Run gofumpt and fail if formatting changes"
	@echo "  make lint         Run golangci-lint"
	@echo "  make tidy         Run go mod tidy and fail if changes occur"

infra-up:
	docker compose -f infra/docker-compose.yml up -d

infra-down:
	docker compose -f infra/docker-compose.yml down

dev:
	set -a && . ./.env && set +a && go run ./cmd/api

test:
	go test ./...

fmt-check:
	gofumpt -w -l .
	@git diff --exit-code

lint:
	golangci-lint run ./...

tidy:
	go mod tidy
	@git diff --exit-code