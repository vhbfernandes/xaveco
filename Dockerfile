FROM golang:1.16-alpine as build

ENV APP_DIR $GOPATH/src/github.com/hellofreshdevtests/vhbfernandes-sre-test

RUN apk add --update --no-cache git && \
    mkdir -p $APP_DIR

ADD . $APP_DIR

WORKDIR $APP_DIR
RUN go get
RUN CGO_ENABLED=0 go build -o /tmp/simpleapi main.go

FROM alpine:3.14

COPY --from=build /tmp/simpleapi /simpleapi

ENTRYPOINT "/simpleapi"