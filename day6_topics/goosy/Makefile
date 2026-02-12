-include .env

DB_DRIVER?=postgres
DB_USER?=postgres
DB_PASS?=postgres
DB_HOST?=localhost
DB_PORT?=5432
DB_NAME?=appdb
DB_SSLMODE?=disable

MIGRATIONS_DIR=./db/migrations


DB_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)


.PHONY: up down status reset create down-to up1

up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" up


down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" down

status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" status


reset:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" reset


up1:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" up-by-one

down-to:
	@test -n "$(version)" || (echo "usage: make down-to version=20260211051218" && exit 1)
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_URL)" down-to $(version)

db-up:
	docker compse up -d

db-down:
	docker compose down

db-logs:
	docker compose logs -f db

psql:
	psql "$(DB_URL)"

fresh: reset up

run:
	go run ./cmd/app

test:
	go test ./...

create:
	@test -n "$(name)" || (echo "usage: make create name=create_users_table" && exit 1)
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

