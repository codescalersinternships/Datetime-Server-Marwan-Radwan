# Datetime-Server-Marwan-Radwan

Am implementation of a basic HTTP server in Go that returns the current date and time. The server is implemented using net/http library and Gin framework.

## Usage

Use the Makefile to build and run the project.

To build the HTTP server

```
make build http
```

to build the Gin server

```
make build gin
```

Run using docker compose:

```bash
docker-compose -f docker-compose.yml up --build -d
```

Access the (net/http) server endpoint:

```
curl http://localhost:8081/datetime
```

Access the Gin server endpoint:

```
curl http://localhost:8080/datetime
```

## Testing

Run the tests using Go's testing package.

```
make test
```
