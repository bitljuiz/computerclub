FROM golang:1.22.2
WORKDIR /app
COPY . .
RUN go build -o /app/main ./cmd/main.go
ENTRYPOINT ["/app/main"]