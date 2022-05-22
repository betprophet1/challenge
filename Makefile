docker-build-local:
	docker build -f build/Dockerfile.local -t prophet/wager-local .
docker-build-dev:
	docker build -f build/Dockerfile.dev -t prophet/wager-dev .
docker-build-prod:
	docker build -f build/Dockerfile.prod -t prophet/wager-prod .
server:
	go run cmd/main.go
docker-compose-up:
	docker-compose --file=build/docker-compose.yml up