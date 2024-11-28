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
deps: deps-accounts deps-users deps-shared

deps-accounts:
	@go mod tidy -C ./services/accounts
	@wire ./services/accounts/cmd/accounts

deps-users:
	@go mod tidy -C ./services/users
	@wire ./services/users/cmd/users

deps-shared:
	@go mod tidy -C ./shared

# update mocks
mocks: mocks-users mocks-accounts

mocks-users:
	@cd ./services/users && \
	mockery && \
	cd ../..

mocks-accounts:
	@cd ./services/accounts && \
	mockery && \
	cd ../..

### ci commands

# build services
go-build: go-build-accounts go-build-users

go-build-accounts:
	@go build -C ./services/accounts -o cmd/accounts .

go-build-users:
	@go build -C ./services/users -o ./cmd/users .

# run tests
test: test-accounts test-users

test-accounts:
	@echo "Running accounts service tests..."
	@gotestsum --format-icons hivis ./services/accounts/...

test-users:
	@echo "\nRunning users service tests..."
	@gotestsum --format-icons hivis ./services/users/...

# run linters
lint: lint-accounts lint-users lint-shared

lint-accounts:
	@go vet ./services/accounts/...

lint-users:
	@go vet ./services/users/...

lint-shared:
	@go vet ./shared/...

### Database commands
db-status:
	@echo "accounts service migrations"
	@make run-migration service=accounts cmd=status
	@echo "users service migrations"
	@make run-migration service=users cmd=status

migrate-all:
	make migrate service=accounts
	make migrate service=users

run-migration:
	docker exec transactions-api-$(service)-1 goose -dir "internal/infra/db/migrations" $(cmd)

migrate:
	make run-migration service=$(service) cmd=up

rollback:
	make run-migration service=$(service) cmd=down

migration:
	docker exec transactions-api-$(service)-1 goose -dir "internal/infra/db/migrations" create $(name) sql

