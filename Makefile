dev:
	go run ./cmd/sandbox/main.go 

docker-build:
	docker build -f ./cmd/sandbox/Dockerfile -t app-sandbox .

docker-run:
	docker run -p 127.0.0.1:8080:8080/tcp --name app-sandbox app-sandbox

docker-restart:
	make docker-build
	make docker-run

ecr-login:
	aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/sumeragis-rest

ecr-push:
	docker build -f ./build/Dockerfile -t sandbox .
	docker tag sandbox:latest public.ecr.aws/sumeragis-rest/sandbox:latest
	docker push public.ecr.aws/sumeragis-rest/sandbox:latest