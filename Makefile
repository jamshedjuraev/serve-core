postgres-docker:
	@docker run -d --rm \
		--name backend-master-class \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=postgres \
		-p 5433:5432 \
		postgres:latest

nats-docker:
	@docker run -d --rm \
		--name nats-server \
		-p 4222:4222 \
		-p 6222:6222 \
		-p 8222:8222 \
		nats:latest

stop-docker:
	@docker stop backend-master-class nats-server

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
