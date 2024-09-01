.PHONY: all init build_docker launch_docker down_docker

# Colores
COLOR_RESET = \033[0m
COLOR_GREEN = \033[32m
COLOR_YELLOW = \033[33m
COLOR_BLUE = \033[34m
COLOR_RED = \033[31m

all: init launch

init:
	@echo "$(COLOR_BLUE)Starting project...$(COLOR_RESET)"
	@go get -v ./... && go mod tidy

launch:
	@echo "$(COLOR_GREEN)Launching backend...$(COLOR_RESET)"
	@go run main.go

build_docker:
	@echo "$(COLOR_YELLOW)Compiling Quick QR to Docker...$(COLOR_RESET)"
	@docker compose build
	$(MAKE) launch_docker

launch_docker:
	@echo "$(COLOR_GREEN)Launching Quick QR into Docker...$(COLOR_RESET)"
	@docker compose build
	@docker compose up -d

down_docker:
	@echo "$(COLOR_RED)Stopping and removing Quick QR from Docker...$(COLOR_RESET)"
	@docker compose down -v
