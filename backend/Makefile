include .env
export

.PHONY: run run-worker build tidy migrate-up migrate-down migrate-create \
        docker-up docker-down docker-build docker-logs test lint

run:
	go run ./cmd/api

run-worker:
	go run ./cmd/worker

build:
	go build -o bin/api ./cmd/api
	go build -o bin/worker ./cmd/worker

tidy:
	go mod tidy

migrate-up:
	migrate -path migrations -database "$(DB_URL_LOCAL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL_LOCAL)" -verbose down 1

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

docker-up:
	docker compose -f deploy/docker/docker-compose.yml --env-file .env up -d

docker-down:
	docker compose -f deploy/docker/docker-compose.yml down

docker-build:
	docker compose -f deploy/docker/docker-compose.yml build

docker-logs:
	docker compose -f deploy/docker/docker-compose.yml logs -f api

test:
	go test ./... -v -cover

lint:
	golangci-lint run ./...
