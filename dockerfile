FROM golang:1.17 AS builder
WORKDIR /app
COPY . ./
RUN go mod download
# extldflags and linkmode resolve notfound error for sqlite
RUN  GOOS=linux GOARCH=amd64 go build -o /app/main -a -ldflags '-linkmode external -extldflags "-static"'

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT [ "/app/main"]