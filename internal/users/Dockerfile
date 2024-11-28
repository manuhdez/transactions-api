FROM golang:1.23-alpine AS builder

RUN go install github.com/air-verse/air@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM golang:1.23-alpine

COPY --from=builder /go/bin/air /usr/local/bin/air
COPY --from=builder /go/bin/goose /usr/local/bin/goose

WORKDIR /app

COPY ./internal/users /internal/users
COPY ./shared /shared

# Install dependencies in the users folder
WORKDIR /internal/users
RUN go mod tidy
RUN go mod verify

# Install dependencies in the shared folder
WORKDIR /shared
RUN go mod tidy
RUN go mod verify

# Reset workdir back to normal
WORKDIR /app

COPY ./internal/users/.air.toml /

ENTRYPOINT ["air"]
