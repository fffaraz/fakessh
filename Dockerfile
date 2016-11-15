FROM golang:alpine
MAINTAINER Faraz Fallahi <fffaraz@gmail.com>
EXPOSE 22
ENTRYPOINT ["fakessh"]
RUN \
	apk add --update --no-cache git && \
	go get github.com/fffaraz/fakessh && \
	apk del git pcre expat libcurl libssh2 && \
	rm -rf /go/pkg /go/src /var/cache/apk/*
