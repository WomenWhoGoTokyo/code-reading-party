# From: golang
FROM golang:latest AS stage1

WORKDIR /app
COPY . .

RUN go build -o simple
# マルチ
FROM ubuntu:latest
COPY --from=stage1 /app/simple /app/simple
EXPOSE 8080
CMD ["/app/simple"]