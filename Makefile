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

ginserver:
	go run ./cmd/timeGin/main.go

gin_build:
	go build ./pkg/timeGin/server.go