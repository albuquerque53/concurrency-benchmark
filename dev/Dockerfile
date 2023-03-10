FROM golang:1.19 as base

FROM base as dev

RUN go install github.com/tsliwowicz/go-wrk@latest

WORKDIR /opt/api

CMD ["go", "run", "./cmd/main.go"]
