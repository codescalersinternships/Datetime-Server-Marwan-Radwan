httpserver:
	go run ./cmd/timeHttp/main.go

test:
	go test -v -cover ./...

format:
	go fmt ./...

linter:
	golint ./...


http_build:
	go build ./pkg/timeHttp/server.go