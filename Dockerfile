FROM golang:1.25.3-alpine AS builder
WORKDIR /app

# Enable Go modules and speed up downloads
ENV CGO_ENABLED=0 GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o parse-message ./main

FROM alpine:3.20
WORKDIR /app

# Certificates for outbound HTTPS (Gemini, etc.)
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/parse-message /app/parse-message
COPY resources /app/resources

EXPOSE 50051

CMD ["/app/parse-message"]
