build: deps test vet
	go build -o bin/pd ./command

deps:
	go get ./...

test: deps
	go test -v -race -cover ./...

vet: deps
	go vet ./...

deploy:
	- curl -sL https://git.io/goreleaser | bash

lint:
	golangci-lint run
