http server:
	go run ./cmd/timeHttp/main.go

build http:
	go build ./cmd/timeHttp/main.go

gin server:
	go run ./cmd/timeGin/main.go

build gin:
	go build ./cmd/timeGin/server.go

test:
	go test -v -cover ./...

format:
	go fmt ./...

lint:
	golangci-lint run