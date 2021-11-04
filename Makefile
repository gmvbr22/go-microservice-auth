run:
	go run api/main.go

build:
	go build -o go_app api/main.go

test:
	go test ./...

format:
	go fmt ./... 