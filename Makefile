include .env
export

## Run the application
run:
	go run cmd/api/main.go

migrate-up:
	migrate -path database/migrations -database ${DATABASE_URL} up

migrate-down:
	migrate -path database/migrations -database ${DATABASE_URL} down

migrate-drop:
	migrate -path database/migrations -database ${DATABASE_URL} drop

migrate-down-all:
	migrate -path database/migrations -database ${DATABASE_URL} down -all

migrate-delete-force:
	migrate -path database/migrations -database ${DATABASE_URL} force ${VERSION}

## Run migrations
# pg-migrate:
# 	docker compose --profile tools run --rm migrate up

## Rollback all migrations
# pg-redo:
# 	docker compose --profile tools run --rm migrate down

## Generate db models
gen-models:
	sqlc generate
#sqlboiler psql

test:
	go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./...