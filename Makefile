#!make

vendor:
	docker run --rm -it -v "${PWD}":/app -w /app golang:1.14-alpine go mod vendor
