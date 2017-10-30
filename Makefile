.PHONY: build

build:
	@go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d)" -o createstruct

run: build
	@./createstruct

release: *.go service/*.go *.md
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d)" -a -o createstruct
	docker build -t vikings/createstruct .
	docker push vikings/createstruct