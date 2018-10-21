FROM golang:alpine as build

COPY . $GOPATH/src/github.com/free/jiralert

RUN apk add git && \
    go get $GOPATH/src/github.com/free/jiralert/ && \
    go build -o jiralert $GOPATH/src/github.com/free/jiralert/cmd/jiralert

FROM alpine
COPY --from=build jiralert /usr/local/bin/jiralert

ENTRYPOINT ["/usr/local/bin/jiralert"]
