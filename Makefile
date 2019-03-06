.PHONY: start build run lint clean

GONAME=goWechat

default: build

start:
	@RUN_MODE=prod ./bin/$(GONAME) 

build:
	@go build -o bin/$(GONAME) 

run:
	@./bin/$(GONAME) 

lint:
	@golint

clean:
	@go clean && rm -rf ./bin/$(GONAME) && rm -f go_wechat
