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
