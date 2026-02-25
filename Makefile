.PHONY: help dev infra-up infra-down

help:
	@echo "Targets:"
	@echo "  make infra-up     Start postgres + redis"
	@echo "  make infra-down   Stop postgres + redis"
	@echo "  make dev          Run API locally"

infra-up:
	docker compose -f infra/docker-compose.yml up -d

infra-down:
	docker compose -f infra/docker-compose.yml down

dev:
	set -a && . ./.env && set +a && go run ./cmd/api