.DEFAULT_GOAL:=help
.PHONY: help run build test clean tidy docker-up docker-down dev start migrate-create migrate-up migrate-down apt-install-migrate pacman-install-migrate

# Run the application
run:
	go run cmd/api/main.go

# Build the application
build:
	go build -o bin/api cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Download dependencies
tidy:
	go mod tidy

# Start PostgreSQL container
docker-up:
	docker compose up -d

# Stop PostgreSQL container
docker-down:
	docker compose down

# Install dependencies and run
dev: tidy run

# Full setup: start docker, tidy dependencies, and run
start: docker-up tidy run

# Install the migrate CLI tool (Debian/Ubuntu)
apt-install-migrate: ## Install the migrate CLI tool (apt)
	@echo "Installing migrate CLI..."
	@sudo apt install --no-install-recommends -y postgresql-client curl && \
	curl -L -o /tmp/migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz && \
	sudo tar xzf /tmp/migrate.tar.gz -C /usr/local/bin && \
	sudo chmod +x /usr/local/bin/migrate
	@echo "migrate installation completed"

# Install the migrate CLI tool (Arch Linux)
pacman-install-migrate: ## Install the migrate CLI tool (pacman)
	@echo "Installing migrate CLI..."
	@sudo pacman -S --noconfirm postgresql curl && \
	curl -L -o /tmp/migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz && \
	sudo tar xzf /tmp/migrate.tar.gz -C /usr/local/bin && \
	sudo chmod +x /usr/local/bin/migrate
	@echo "migrate installation completed"

# Create a new migration
migrate-create: ## Create a new migration
	@read -p "Enter migration name: " name; \
	migrate create -dir migrations -ext sql $$name

# Apply all available migrations
migrate-up: ## Apply all available migrations
	@migrate -path migrations -database 'postgres://postgres:postgres@localhost:5432/wappi?sslmode=disable' up

# Revert the last migration
migrate-down: ## Revert the last migration
	@migrate -path migrations -database 'postgres://postgres:postgres@localhost:5432/wappi?sslmode=disable' down

# Show help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | grep ^help -v | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'
