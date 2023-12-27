BASE := ./deployments/docker-compose.yml
ENV := ./deployments/.env.default

BASEF := -f $(BASE) --env-file $(ENV)

test:
	go test ./...

generate:
	sqlc generate
	easyjson ./internal/saver/service/models.go
	go generate ./...

lint:
	gofumpt -w cmd/ internal/
	golangci-lint run cmd/... internal/...

build:
	go build -o l0 cmd/l0/main.go

compose-build:
	docker compose $(BASEF) build

compose-up:
	docker compose $(BASEF) up

compose-down:
	docker compose $(BASEF) down



.PHONY: test generate lint build compose-build compose-up compose-down 
