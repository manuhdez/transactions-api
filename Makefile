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
deps: deps-accounts deps-transactions

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

# go commands
#
# Deps
tidy: tidy-accounts tidy-transactions

tidy-accounts:
	@cd internal/accounts && go mod tidy && cd -

tidy-transactions:
	@cd internal/transactions && go mod tidy && cd -

# Testing
test: test-accounts test-transactions

test-accounts:
	@go test ./internal/accounts/...

test-transactions:
	@go test ./internal/transactions/...

# Code linting
lint: lint-accounts lint-transactions

lint-accounts:
	@go vet ./internal/accounts/...

lint-transactions:
	@go vet ./internal/transactions/...

# Database migrations
#
.PHONY: goose
goose:
	GOOSE_DRIVER=mysql goose mysql $(db) $(cmd)

goose-transactions:
	make goose db="%user:%password@tcp(%transactions_db)/%transactions-api" cmd=$(cmd)

migration:
	docker exec transactions-api-$(service)-1 goose -dir "db/migrations" $(cmd)

migrate: migrate-transactions migrate-accounts

migrate-accounts:
	@echo "Running accounts service migrations 🚀" && \
	make migration service=accounts cmd=up
migrate-transactions:
	@echo "Running transactions service migrations 🚀" && \
	make migration service=transactions cmd=up

migration-create:
	docker exec transactions-api-$(service)-1 goose -dir "db/migrations" create $(name) sql
