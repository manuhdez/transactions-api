# docker compose
build:
	@echo "🛠 Building containers"
	@docker-compose build --no-cache

up: build
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
	@cd internal/accounts && go test ./... && cd -

test-transactions:
	@cd internal/transactions && go test ./... && cd -

# Code linting
lint: lint-accounts lint-transactions

lint-accounts:
	@cd internal/accounts && go vet ./... && cd -

lint-transactions:
	@cd internal/transactions && go vet ./... && cd -
