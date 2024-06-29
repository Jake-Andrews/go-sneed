.PHONY: local docker run down-local

local:
	echo "Running postgres"
	docker compose -f docker-compose.local.yml up --build

#FILES := $(shell docker ps -a -f name=postgres)
down-local:
	docker compose -f docker-compose.local.yml down -v
	#docker stop $(FILES)
	#docker rm $(FILES)

run:
	go run cmd/main.go

