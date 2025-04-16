FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o cmd/api/comment-moderation ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/cmd/api/comment-moderation /app/
EXPOSE 8080
CMD ["./comment-moderation"]