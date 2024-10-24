FROM golang:1.23

RUN go install github.com/cespare/reflex@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY reflex.conf /
COPY start.sh /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
