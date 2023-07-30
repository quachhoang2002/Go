include .env
export

## Run the application
run:
	go run cmd/api/main.go

## Run migrations
pg-migrate:
	docker compose --profile tools run --rm migrate up

## Rollback all migrations
pg-redo:
	docker compose --profile tools run --rm migrate down

## Generate db models
gen-models:
	sqlboiler psql

test:
	go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./...