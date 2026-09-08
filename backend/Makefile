include .env
export

.PHONY: run run-worker seed-admin build tidy migrate-up migrate-down migrate-create \
        docker-up docker-down docker-build docker-logs test lint

run:
	go run ./cmd/api

run-worker:
	go run ./cmd/worker

# Bootstrap (or promote) the first admin account. Requires ADMIN_EMAIL,
# ADMIN_USERNAME, ADMIN_PASSWORD in .env (ADMIN_DISPLAY_NAME optional).
seed-admin:
	go run ./cmd/seed

build:
	go build -o bin/api ./cmd/api
	go build -o bin/worker ./cmd/worker

tidy:
	go mod tidy

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down 1

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

docker-up:
	docker compose -f deploy/docker/docker-compose.yml --env-file .env up -d

# Also build+run api & worker as containers (profile "app"), instead of
# `go run` natively. Use this if you don't want Go installed on the host.
docker-up-app:
	docker compose -f deploy/docker/docker-compose.yml --env-file .env --profile app up -d --build

docker-down:
	docker compose -f deploy/docker/docker-compose.yml --profile app down

docker-build:
	docker compose -f deploy/docker/docker-compose.yml build

docker-logs:
	docker compose -f deploy/docker/docker-compose.yml logs -f api

test:
	go test ./... -v -cover

lint:
	golangci-lint run ./...
