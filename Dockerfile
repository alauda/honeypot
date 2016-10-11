FROM golang:1.6

MAINTAINER bin liu

ENTRYPOINT ["/honeypot/honeypot"]
EXPOSE 8080

ADD . /go/src/honeypot
WORKDIR /go/src
RUN go build -v -o /honeypot/honeypot honeypot
RUN chmod +x /honeypot/honeypot

RUN rm -rf /go/src/honeypot && mkdir -p /honeypot/logs

WORKDIR /honeypot
