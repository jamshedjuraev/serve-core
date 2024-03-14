# Инструкция по запуску сервиса локально:
#	1) make up
#	2) make nats-url
#	3) make nats-kv
# 	4) make run

nats-url:
	@echo "nats://localhost:4222"

up:
	@docker-compose up -d

down:
	@docker-compose down

run:
	go run cmd/app/main.go

nats-kv:
	@nats kv add sharedConfig
	@nats kv put sharedConfig SERVER_PORT :8080
	@nats kv put sharedConfig DATABASE_DSN postgres://postgres:postgres@localhost:5433/postgres

.PHONY: postgres-docker nats-docker stop-docker nats-url run
