.PHONY: all init build_docker launch_docker down_docker

# Colores
COLOR_RESET = \033[0m
COLOR_GREEN = \033[32m
COLOR_YELLOW = \033[33m
COLOR_BLUE = \033[34m
COLOR_RED = \033[31m

all: help

init: ## Download dependencies
	@echo "$(COLOR_BLUE)Starting project...$(COLOR_RESET)"
	@go get -v ./... && go mod tidy

launch: ## Start docker local
	@echo "$(COLOR_GREEN)Launching backend...$(COLOR_RESET)"
	@go run main.go

build_docker: ## Build docker api
	@echo "$(COLOR_YELLOW)Compiling Quick QR to Docker...$(COLOR_RESET)"
	@docker compose build
	$(MAKE) launch_docker

launch_docker: ## Start docker api
	@echo "$(COLOR_GREEN)Launching Quick QR into Docker...$(COLOR_RESET)"
	@docker compose up -d

down_docker: ## Stops docker api
	@echo "$(COLOR_RED)Stopping and removing Quick QR from Docker...$(COLOR_RESET)"
	@docker compose down -v

test: ## Run the tests
	@echo "$(COLOR_RED)Starting tests...$(COLOR_RESET)"
	@go test ./...

help: ## Shows the help
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
        awk 'BEGIN {FS = ":.*?## "}; {printf "$(COLOR_YELLOW)make $(COLOR_BLUE)%-20s$(COLOR_BLUE) %s\n", $$1, $$2}'
	@echo ''