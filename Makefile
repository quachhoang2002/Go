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

proto:
	protoc --proto_path=proto  --go_out=pb  --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

## Generate db models
gen-models:
	sqlc generate
	
test:
	go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./...

.PHONY: run migrate-up migrate-down migrate-drop migrate-down-all migrate-delete-force gen-models test proto