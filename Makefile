up:
	docker-compose up --build -d

up-recreate:
	docker-compose down -v && make up

stop:
	docker-compose stop

docker-build:
	docker build -f ./cmd/echo/Dockerfile -t app-echo .

migrate-up:
	migrate -database mysql://docker:docker@tcp(localhost:3306)/general -path ./schema/mysql/ddl up

dev:
	go run ./cmd/sandbox/main.go 

app-build:
	docker build -f ./cmd/sandbox/Dockerfile -t app-sandbox .

app-run:
	docker run -p 127.0.0.1:8080:8080/tcp --name app-sandbox app-sandbox

ecr-login:
	aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/sumeragis-rest

ecr-push:
	docker build -f ./build/Dockerfile -t sandbox .
	docker tag sandbox:latest public.ecr.aws/sumeragis-rest/sandbox:latest
	docker push public.ecr.aws/sumeragis-rest/sandbox:latest