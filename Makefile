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
	@docker ps --format "📦 {{.ID}} - {{.Image}} ⏱  {{.Status}}"

# dependencies
# generate dependency tree
deps: deps-accounts

deps-accounts: 
	@cd internal/accounts && \
	wire && \
	cd - && \
	make tidy-accounts

# go commands
#
# Deps
tidy: tidy-accounts

tidy-accounts:
	@cd internal/accounts && go mod tidy && cd -

# Testing
test: test-accounts

test-accounts:
	@cd internal/accounts && go test ./... && cd -
