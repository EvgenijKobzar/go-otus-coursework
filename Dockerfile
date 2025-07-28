FROM golang:1.24.2-alpine as builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc musl-dev

COPY ["app/go.mod", "app/go.sum", "./"]

RUN go mod download

COPY app ./

RUN go build -o ./bin/app cmd/server/main.go

FROM alpine as runner

COPY --from=builder /usr/local/src/bin/app /
COPY static /static
COPY .env.prod /.env
EXPOSE 8080

CMD ["/app"]