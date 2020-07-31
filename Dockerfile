FROM golang:1.14 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build the application
RUN go build -o server greeter_server/main.go

FROM debian:bullseye-slim AS runner

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

COPY --from=builder /app/server /app/server

ENTRYPOINT ["/app/server"]
