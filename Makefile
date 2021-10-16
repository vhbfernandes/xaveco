test:
	docker compose up -d mongo && MONGODB_CONNECTION_STRING="mongodb://root:root@mongo:27017/test" MONGODB_DATABASE=test go test ./pkg/repository

up:
	docker compose up --build

build:
	docker build -t victorhbfernandes/xaveco:latest . && docker push victorhbfernandes/xaveco:latest
