# Variables
LATEST_COMMIT := $$(git rev-parse HEAD)
VERSION ?= latest
VER ?= latest

run-local: ## Run via `go run`
	@DATABASE_URL="postgres://postgres:123456@localhost:5432/testdb?sslmode=disable" \
	CONSOLE_SERVER_ADDRESS=:8088 \
	go run -ldflags "-X main.buildTag=`date -u +%Y%m%d.%H%M%S` -$(LATEST_COMMIT)" cmd/awesomeProject/main.go run