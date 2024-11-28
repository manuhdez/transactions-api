.PHONY: migration migrate

# docker commands
build:
	@docker-compose build

up:
	@docker-compose up -d

down:
	@docker-compose down

### code commands

# manage dependencies
deps: deps-accounts deps-transactions deps-users deps-shared

deps-accounts:
	@go mod tidy -C ./services/accounts
	@wire ./services/accounts/cmd/accounts

deps-transactions:
	@go mod tidy -C ./services/transactions
	@wire ./services/transactions

deps-users:
	@go mod tidy -C ./services/users
	@wire ./services/users/cmd/users

deps-shared:
	@go mod tidy -C ./shared

# update mocks
mocks: user-mocks accounts-mocks transactions-mocks

user-mocks:
	@cd ./services/users && \
	mockery && \
	cd ../..

accounts-mocks:
	@cd ./services/accounts && \
	mockery && \
	cd ../..

transactions-mocks:
	@cd ./services/transactions && \
	mockery && \
	cd ../..

### ci commands

# build services
go-build: go-build-accounts go-build-transactions go-build-users

go-build-accounts:
	@go build -C ./services/accounts -o cmd/accounts .

go-build-transactions:
	@go build -C ./services/transactions -o ./cmd/transactions .

go-build-users:
	@go build -C ./services/users -o ./cmd/users .

# run tests
test: test-accounts test-transactions test-users

test-accounts:
	@echo "Running accounts service tests..."
	@gotestsum --format-icons hivis ./services/accounts/...

test-transactions:
	@echo "\nRunning transactions service tests..."
	@gotestsum --format-icons hivis ./services/transactions/...

test-users:
	@echo "\nRunning users service tests..."
	@gotestsum --format-icons hivis ./services/users/...

# run linters
lint: lint-accounts lint-transactions lint-users lint-shared

lint-accounts:
	@go vet ./services/accounts/...

lint-transactions:
	@go vet ./services/transactions/...

lint-users:
	@go vet ./services/users/...

lint-shared:
	@go vet ./shared/...



### Database commands
db-status:
	@echo "accounts service migrations"
	@docker exec transactions-api-accounts-1 goose -dir "internal/infra/db/migrations" status
	@echo "transactions service migrations"
	@docker exec transactions-api-transactions-1 goose -dir "db/migrations" status
	@echo "users service migrations"
	@docker exec transactions-api-users-1 goose -dir "internal/infra/db/migrations" status

migrate-all:
	make migrate service=accounts
	make migrate service=transactions
	make migrate service=users

run-migration:
	docker exec transactions-api-$(service)-1 goose -dir "internal/infra/db/migrations" $(cmd)

migrate:
	make run-migration service=$(service) cmd=up

rollback:
	make run-migration service=$(service) cmd=down

migration:
	docker exec transactions-api-$(service)-1 goose -dir "internal/infra/db/migrations" create $(name) sql

