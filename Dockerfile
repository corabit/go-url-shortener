FROM golang:1.17 AS builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o /app/main -a

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT [ "/app/main"]