# Compile Go binary.
FROM golang:1.24.4-alpine AS builder
RUN apk add --no-cache git
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /workspace/cmd/server

# Disable CGO as we don't need to call C libs, improve portability.
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN go build -ldflags="-s -w" -o currency-app

# Run in tiny alpine image.
FROM alpine:latest
RUN apk add --no-cache ca-certificates # Add root CA, we're calling API over https..
WORKDIR /app
COPY --from=builder /workspace/cmd/server/currency-app .
EXPOSE 8080
ENTRYPOINT ["./currency-app"]