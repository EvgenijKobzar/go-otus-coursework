FROM golang:1.24.2-alpine AS goose-builder
RUN apk add --no-cache git
RUN go install github.com/pressly/goose/v3/cmd/goose@latest


FROM alpine AS runner

RUN apk --no-cache add bash git make gcc musl-dev postgresql-client

COPY --from=goose-builder /go/bin/goose /usr/local/bin/goose
RUN goose -version
COPY ../../app/internal/migrations /migrations