FROM golang:alpine AS builder
RUN apk add --update git && go get github.com/fffaraz/fakessh

FROM alpine:latest
COPY --from=builder /go/bin/fakessh /usr/local/bin
EXPOSE 22
ENTRYPOINT ["fakessh"]
