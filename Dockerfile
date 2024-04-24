FROM golang:1.21.6 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o out ./cmd/mygo

FROM alpine

WORKDIR /app

RUN mkdir config && mkdir dict

COPY --from=builder /app/out .
COPY --from=builder /app/config/config.yaml ./config/
COPY --from=builder /app/dict ./dict/

ENV GIN_MODE=release

EXPOSE 80

ENTRYPOINT ["./out"]