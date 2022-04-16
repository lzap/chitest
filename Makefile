run:
	go run main.go

build:
	go build

models: sqlboiler.toml
	sqlboiler sqlite3 --wipe -o pkg/models

migrate:
	sqlite3 devel.sqlite3 < cmd/migrate/schema.sql

.PHONY: build models migrate
