v_file := VERSION
VERSION := $(shell cat ${v_file})


help:
	@echo "  docker - build docker image"
	@echo "  build - compile app"


docker:
	env GOOS=linux GOARCH=amd64 go build -o goapi-linux goapi.go
	docker build -t webapi:${VERSION} .

build:
	go build -o goapi-darwin goapi.go 

build_linux:
	env GOOS=linux GOARCH=amd64 go build -o goapi-linux goapi.go
