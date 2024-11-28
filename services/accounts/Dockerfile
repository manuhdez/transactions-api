FROM golang:1.23-alpine AS builder

RUN go install github.com/air-verse/air@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM golang:1.23-alpine

COPY --from=builder /go/bin/air /usr/local/bin/air
COPY --from=builder /go/bin/goose /usr/local/bin/goose

WORKDIR /app

COPY ./services/accounts /services/accounts
COPY ./shared /shared

# Install dependencies in the accounts folder
WORKDIR /services/accounts
RUN go mod tidy
RUN go mod verify

# Install dependencies in the shared folder
WORKDIR /shared
RUN go mod tidy
RUN go mod verify

# Reset workdir back to normal
WORKDIR /app

COPY ./services/accounts/.air.toml /

ENTRYPOINT ["air"]
