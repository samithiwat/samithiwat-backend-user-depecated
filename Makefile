test:
	go test ./src/test

server:
	go run ./src/.

compose-up:
	docker-compose --env-file .env.dev up -d

compose-down:
	docker-compose --env-file .env.dev down
