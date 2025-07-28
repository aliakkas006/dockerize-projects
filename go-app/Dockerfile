FROM golang:1.24.5-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

# Build a statically linked binary for Alpine (musl)
RUN go build -ldflags="-s -w" -o server

EXPOSE 8080

CMD ["./server"]
