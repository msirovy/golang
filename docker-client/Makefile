build:
	GOOS=linux GOARCH=amd64 go build .

release:
	scp ./docker-cli n2-prod.freelo.io:/root/docker-cli
	