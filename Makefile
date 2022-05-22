docker-build-local:
	docker build -f build/Dockerfile.local -t prophet/wager-local .
docker-build-dev:
	docker build -f build/Dockerfile.dev -t prophet/wager-dev .
docker-build-prod:
	docker build -f build/Dockerfile.prod -t prophet/wager-prod .
run-docker:
	docker run -d --name wager -p 8080:8080 -e MYSQL_USER='wager' \
	-e MYSQL_HOST='localhost' \
	-e MYSQL_PORT='3306' \
	-e MYSQL_PASSWORD='123456' \
	-e MYSQL_DATABASE='wager' prophet/wager-local
server:
	go run cmd/main.go
docker-compose-up:
	docker-compose up