FROM golang:1.24.5-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

# Build a statically linked binary for Alpine (musl)
RUN go build -ldflags="-s -w" -o server

EXPOSE 8080

CMD ["./server"]




# # Build stage
# FROM golang:1.24.5-alpine AS builder
# WORKDIR /app
# COPY . .
# RUN go build -o main .

# # Runtime stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /app/main .
# CMD ["./main"]