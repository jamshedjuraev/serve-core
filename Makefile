run-docker:
	@docker run -d --rm \
		--name serve-core \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=postgres \
		-p 5433:5432 \
		postgres:16-alpine

run:
	go run main.go

.PHONY: run-docker run
