 export GO111MODULE=on
# GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
# GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)


default: build

build: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o moviedb-api 