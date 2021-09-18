FROM golang:1.15-alpine as build

WORKDIR /go/src/github.com/Nikita99100/Fibonacci

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
  -o bin/fibonacci \
  ./cmd

FROM alpine:3.12

COPY --from=build /go/src/github.com/Nikita99100/Fibonacci/bin/fibonacci /usr/local/bin/fibonacci
COPY --from=build /go/src/github.com/Nikita99100/Fibonacci/config/configs.toml /etc/fibonacci.toml

CMD ["/usr/local/bin/fibonacci", "-config", "/etc/fibonacci.toml"]