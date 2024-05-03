# Инструкция по запуску сервиса локально:
#	1) make up			(поднять бд в докере)
#	2) make set-env		(назначить CONFIG_PATH в env)
#	3) make run


set-env:
	@export CONFIG_PATH=./config/local.yaml

up:
	@docker-compose up -d

down:
	@docker-compose down

run:
	go run cmd/app/main.go

.PHONY: set-env up down run
