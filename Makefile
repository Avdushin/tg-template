# Makefile
APP_NAME=bot
DOCKER_IMAGE=mybot:latest
COMPOSE_FILE=docker-compose.yml

# MODE может быть "local", "docker", "compose"
MODE ?= local

.PHONY: all run build air clean

# ----------------------------
# Основной таргет run
# ----------------------------
run: build
ifeq ($(MODE),local)
	@echo "Running $(APP_NAME) locally..."
	./bin/$(APP_NAME)
else ifeq ($(MODE),docker)
	@echo "Running $(APP_NAME) in Docker..."
	docker build -t $(DOCKER_IMAGE) .
	docker run --env-file .env -it --rm $(DOCKER_IMAGE)
else ifeq ($(MODE),compose)
	@echo "Running $(APP_NAME) with docker-compose..."
	docker compose -f $(COMPOSE_FILE) up --build
else
	$(error Unknown MODE '$(MODE)'. Use MODE=local|docker|compose)
endif

# ----------------------------
# Локальный билд
# ----------------------------
build:
	@echo "Building $(APP_NAME)..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME) ./cmd/bot

# ----------------------------
# Режим live-reload через air
# ----------------------------
air:
	@echo "Running $(APP_NAME) with live-reload..."
	air

# ----------------------------
# Очистка бинарей и временных файлов
# ----------------------------
clean:
	@echo "Cleaning..."
	rm -f bin/$(APP_NAME)
	rm -rf tmp
