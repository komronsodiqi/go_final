build:
    go build -o cmd/api/comment-moderation ./cmd/api

run: build
    ./cmd/api/comment-moderation

docker-build:
    docker build -t comment-moderation-service .

docker-run: docker-build
    docker run -p 8080:8080 comment-moderation-service