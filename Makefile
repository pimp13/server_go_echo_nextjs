include .env
export

BINARY_NAME=server

BIN_DIR=bin

SRC_DIR=cmd/api/main.go

TEST_DIR=./...

DB_STRING ?= $(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

MIGRATION_DIR ?= cmd/migrations

.PHONY: migrate-up migrate-down migrate-status create-migration all build clean start restart test deps

migrate-up:
	@goose -dir $(MIGRATION_DIR) mysql "$(DB_STRING)" up

migrate-down:
	@goose -dir $(MIGRATION_DIR) mysql "$(DB_STRING)" down

migrate-status:
	@goose -dir $(MIGRATION_DIR) mysql "$(DB_STRING)" status

create-migration:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME parameter is required"; \
		echo "Usage: make create-migration NAME=description_of_migration"; \
		exit 1; \
	fi
	@goose -dir $(MIGRATION_DIR) create $(NAME) sql


migrate-reset:
	@goose -dir $(MIGRATION_DIR) mysql "$(DB_STRING)" reset


migrate-to:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION parameter is required"; \
		exit 1; \
	fi
	@goose -dir $(MIGRATION_DIR) mysql "$(DB_STRING)" goto $(VERSION)

all: build

build:
	@go build -o $(BIN_DIR)/$(BINARY_NAME) $(SRC_DIR)

clean:
	@rm -rf $(BIN_DIR)

start:
	@./$(BIN_DIR)/$(BINARY_NAME)

restart: build
	@./$(BIN_DIR)/$(BINARY_NAME)

test:
	@go test $(TEST_DIR)

deps:
	@go mod tidy

update-deps:
	@go get -u ./...
	@go mod tidy
