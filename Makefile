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
	@cd internal/accounts/cmd/accounts && \
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

# generate mocks
mocks: user-mocks accounts-mocks transactions-mocks

user-mocks:
	@cd ./internal/users && \
	mockery && \
	cd ../..

accounts-mocks:
	@cd ./internal/accounts && \
	mockery && \
	cd ../..

transactions-mocks:
	@cd ./internal/transactions && \
	mockery && \
	cd ../..

### go commands

# Build
go-build: go-build-accounts go-build-transactions go-build-users

go-build-accounts:
	@go build -C ./internal/accounts -o cmd/accounts .

go-build-transactions:
	@go build -C ./internal/transactions -o ./cmd/transactions .

go-build-users:
	@go build -C ./internal/users -o ./cmd/users .

# Deps
tidy: tidy-accounts tidy-transactions tidy-users tidy-shared

tidy-accounts:
	@cd internal/accounts && go mod tidy && cd -

tidy-transactions:
	@cd internal/transactions && go mod tidy && cd -

tidy-users:
	@cd internal/users && go mod tidy && cd -

tidy-shared:
	@cd shared && go mod tidy && cd -

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
lint: lint-accounts lint-transactions lint-users lint-shared

lint-accounts:
	@go vet ./internal/accounts/...

lint-transactions:
	@go vet ./internal/transactions/...

lint-users:
	@go vet ./internal/users/...

lint-shared:
	@go vet ./shared/...

# Database commands
#
db-status:
	@echo "Accounts service migrations"
	@docker exec transactions-api-accounts-1 goose -dir "db/migrations" status
	@echo "\nTransactions service migrations"
	@docker exec transactions-api-transactions-1 goose -dir "db/migrations" status
	@echo "\nUsers service migrations"
	@docker exec transactions-api-users-1 goose -dir "db/migrations" status

migrate-all:
	@echo "Running all service migrations 🚀" && \
	make migrate service=accounts && \
	make migrate service=transactions && \
	make migrate service=users

run-migration:
	docker exec transactions-api-$(service)-1 goose -dir "db/migrations" $(cmd)

migrate:
	@echo "Running $(service) service migrations 🚀" && \
	make run-migration service=$(service) cmd=up

rollback:
	@echo "Rolling back $(service) service migrations 🚀" && \
	make run-migration service=$(service) cmd=down

migration:
	docker exec transactions-api-$(service)-1 goose -dir "db/migrations" create $(name) sql

.PHONY: migration migrate
