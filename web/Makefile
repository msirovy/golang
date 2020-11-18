v_file := VERSION
VERSION := $(shell cat ${v_file})


help:
	@echo "  docker - build docker image"
	@echo "  build - compile app"


docker:
	env GOOS=linux GOARCH=amd64 go build -o webapi-linux webapi.go
	docker build -t webapi:${VERSION} .

build:
	go build -o webapi-darwin  webapi.go 

build_linux:
	env GOOS=linux GOARCH=amd64 go build -o webapi-linux webapi.go
