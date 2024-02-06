rundocker:
	@docker run -d --rm \
		--name backend-master-class \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=postgres \
		-p 5433:5432 \
		postgres:latest

# postgres://POSTGRES_USER:POSTGRES_PASSWORD@localhost:5433/dbname?sslmode=disable
migrateup:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable" -verbose down

.PHONY: rundocker migrateup migratedown
