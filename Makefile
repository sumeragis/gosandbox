.PHONY: dev
dev:
	go run ./cmd/sandbox/main.go 

.PHONY: up
up:
	docker-compose up --build -d

.PHONY: up-recreate
up-recreate:
	docker-compose down -v && make up

.PHONY: stop
stop:
	docker-compose stop

.PHONY: docker-build
docker-build:
	docker build -f ./cmd/echo/Dockerfile -t app-echo .

.PHONY: migrate-up
migrate-up:
	migrate -database "mysql://docker:docker@tcp(localhost:3306)/general" -path ./schema/mysql/ddl up

.PHONY: seed
seed:
	bash ./script/seed_fixture_to_mysql.sh

.PHONY: app-build
app-build:
	docker build -f ./cmd/sandbox/Dockerfile -t app-sandbox .

.PHONY: app-run
app-run:
	docker run -p 127.0.0.1:8080:8080/tcp --name app-sandbox app-sandbox

.PHONY: ecr-login
ecr-login:
	aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/sumeragis-rest

.PHONY: ecr-push
ecr-push:
	docker build -f ./build/Dockerfile -t sandbox .
	docker tag sandbox:latest public.ecr.aws/sumeragis-rest/sandbox:latest
	docker push public.ecr.aws/sumeragis-rest/sandbox:latest