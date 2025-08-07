# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

# Compile dari folder cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/app ./cmd/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin/app .

ENTRYPOINT ["./app"]
