setenv: 
	export GO111MODULE=on

build: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o moviedb-api 