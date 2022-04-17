run:
	go run main.go

build:
	go build

models: sqlboiler.toml
	sqlboiler sqlite3 --wipe -o pkg/models

migrate:
	sqlite3 devel.db < cmd/migrate/schema.sql

.PHONY: build models migrate
