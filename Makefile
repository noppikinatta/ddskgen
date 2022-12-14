.PHONY: gen run test test-cov build

gen:
	go generate ./...

run:
	go run app/main.go

test:
	go test -v ./...

test-cov:
	go test -cover -coverprofile=cover.out -v ./... && go tool cover -html=cover.out

build:
	go build -o=release/ddskgen app/main.go