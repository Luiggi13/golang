.PHONY: all init build_docker launch_docker down_docker
all: init launch

init:
	@echo "Starting project..."
	go get -v ./... && go mod tidy

launch:
	@echo "Launching backend..."
	go run main.go

build_docker:
	@echo "Compiling Quick QR to Docker"
	docker compose build
	$(MAKE) launch_docker

launch_docker:
	@echo "Launch Quick QR into Docker"
	docker compose build
	docker compose up -d

down_docker:
	@echo "Stopping and removing Quick QR to Docker"
	docker compose down -v
