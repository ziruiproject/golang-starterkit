ENV_FILE := .env

include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

up:
	@echo "Starting database..."
	docker-compose --env-file $(ENV_FILE) up -d --force-recreate --remove-orphans

down:
	@echo "Stopping and removing containers..."
	docker-compose --env-file $(ENV_FILE) down --volumes --remove-orphans

restart: down up

test: restart
	sleep 5
	go test -v ./tests/... | tc
