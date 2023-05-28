FROM golang:1.17.1-alpine3.14 as builder
WORKDIR /app
COPY . .

ENV CGO_ENABLED=0

RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -mod=vendor -o ./main ./cmd/app/main.go
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -mod=vendor -o ./migrate ./cmd/migrations/migrate.go


FROM alpine:3.14 as example-template

COPY --from=builder /app/main    /app/main
COPY --from=builder /app/migrate /app/migrate
COPY entrypoint.sh /app/entrypoint.sh
COPY migrations /migrations

RUN chmod +x /app/entrypoint.sh
EXPOSE 8080

CMD ["/app/entrypoint.sh"]
