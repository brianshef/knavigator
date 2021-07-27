# VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")
MAINFILE := cmd/app/main.go
SERVERFILE := web/app/main.go

install:
	go install ${MAINFILE}

test:
	golint ./...
	go test ./...

serve:
	go run ${SERVERFILE}

run:
	go run ${MAINFILE}

run-image:
	docker run -dp 3000:3000 bshef/knavigator
	@echo "Server available at http://localhost:3000"

build:
	mkdir -p bin
	go build -o bin/${PROJECTNAME} ${SERVERFILE}

build-image:
	docker build -t bshef/knavigator .

push-image:
	docker push bshef/knavigator:latest
	docker push bshef/knavigator:${BUILD}

build-all:
	mkdir -p bin
	@echo "Compiling build ${BUILD} for every OS and Platform"
	GOOS=windows GOARCH=amd64 go build -o bin/${PROJECTNAME}-win-x64.exe ${SERVERFILE}
	GOOS=windows GOARCH=386 go build -o bin/${PROJECTNAME}-win-x86.exe ${SERVERFILE}
	GOOS=linux GOARCH=arm go build -o bin/${PROJECTNAME}-linux-arm ${SERVERFILE}
	GOOS=linux GOARCH=386 go build -o bin/${PROJECTNAME}-linux-x86 ${SERVERFILE}
	GOOS=linux GOARCH=arm64 go build -o bin/${PROJECTNAME}-linux-x64 ${SERVERFILE}
	GOOS=darwin GOARCH=amd64 go build -o bin/${PROJECTNAME}-mac-x64 ${SERVERFILE}

all: install test build-all

.PHONY: all install test run run-image build build-image push-image build-all
