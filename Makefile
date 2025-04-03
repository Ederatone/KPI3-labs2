test:
	go test ./...

build:
	go build -o bin/main ./cmd/example/main.go

run:
	go run ./cmd/example/main.go -e "+ 2 3"