run-docker:
	@docker run -d --rm \
		--name backend-master-class \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=postgres \
		-p 5433:5432 \
		postgres:latest

run:
	go run cmd/app/main.go

stop-docker:
	@docker stop backend-master-class

.PHONY: run-docker run stop-docker
