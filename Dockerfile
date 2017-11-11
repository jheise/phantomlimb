FROM jheise/ubuntu-golang:latest

RUN apt-get update && apt-get install -y phantomjs
RUN go get -u github.com/benbjohnson/phantomjs
RUN mkdir -p /go/src/github.com/jheise/phantomlimb

