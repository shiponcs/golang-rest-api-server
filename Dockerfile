FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY .env .env
COPY . .

RUN #go build -o bin/book-store-api-server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/book-store-api-server .

FROM alpine:latest

WORKDIR /app
COPY .env .env
COPY --from=builder /app/bin/book-store-api-server /app/book-store-api-server

EXPOSE 8080

CMD ["/app/book-store-api-server", "serve", "--port", "8080"]
#CMD ["./book-store-api-server"]
