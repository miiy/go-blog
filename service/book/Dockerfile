ARG GO_VERSION=1.17

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /go/build
COPY . .
RUN go mod download
RUN go build -o ./app cmd/server/main.go


FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /go/build/app /app

EXPOSE 8080

ENTRYPOINT ["./app"]
