# docker compose
build:
	@echo "🛠 Building containers"
	@docker-compose build --no-cache

up:
	@echo "▶️  Running application"
	@docker-compose up -d

down:
	@echo "💥 Destroing containers"
	@docker-compose down

start:
	@echo "▶️  Running application"
	@docker-compose start

stop:
	@echo "✋ Stopping application"
	@docker-compose stop

status:
	@docker ps --format "📦 {{.ID}} - {{.Names}} {{.Ports}}"

# dependencies
# generate dependency tree
deps: deps-accounts deps-transactions deps-users

deps-accounts: 
	@cd internal/accounts && \
	wire && \
	cd - && \
	make tidy-accounts

deps-transactions:
	@cd internal/transactions && \
	wire && \
	cd - && \
	make tidy-transactions

deps-users:
	@cd internal/users && \
	wire && \
	cd - && \
	make tidy-users

# go commands
#
# Deps
tidy: tidy-accounts tidy-transactions tidy-users

tidy-accounts:
	@cd internal/accounts && go mod tidy && cd -

tidy-transactions:
	@cd internal/transactions && go mod tidy && cd -

tidy-users:
	@cd internal/users && go mod tidy && cd -

# Testing
test: test-accounts test-transactions test-users

test-accounts:
	@echo "Running accounts service tests..."
	@gotestsum --format-icons hivis ./internal/accounts/...

test-transactions:
	@echo "\nRunning transactions service tests..."
	@gotestsum --format-icons hivis ./internal/transactions/...

test-users:
	@echo "\nRunning users service tests..."
	@gotestsum --format-icons hivis ./internal/users/...

# Code linting
lint: lint-accounts lint-transactions

lint-accounts:
	@go vet ./internal/accounts/...

lint-transactions:
	@go vet ./internal/transactions/...

# Database commands
#

db-status:
	@echo "Accounts service migrations"
	@docker exec transactions-api-accounts-1 goose -dir "db/migrations" status
	@echo "\nTransactions service migrations"
	@docker exec transactions-api-transactions-1 goose -dir "db/migrations" status
	@echo "\nUsers service migrations"
	@docker exec transactions-api-users-1 goose -dir "db/migrations" status

migrate: migrate-transactions migrate-accounts migrate-users

migration:
	docker exec transactions-api-$(service)-1 goose -dir "db/migrations" $(cmd)

migrate-accounts:
	@echo "Running accounts service migrations 🚀" && \
	make migration service=accounts cmd=up
migrate-transactions:
	@echo "Running transactions service migrations 🚀" && \
	make migration service=transactions cmd=up
migrate-users:
	@echo "Running users service migrations 🚀" && \
	make migration service=users cmd=up

migration-create:
	docker exec transactions-api-$(service)-1 goose -dir "db/migrations" create $(name) sql

.PHONY: migration migrate
