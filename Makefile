#!make
DOCKER_CMD:=docker run --rm -it -v "${PWD}":/app --env-file=.env -w /app golang:1.14-alpine

vendor:
	${DOCKER_CMD} go mod vendor

run:
	${DOCKER_CMD} go run .

test:
	${DOCKER_CMD} go test -v -race ./..