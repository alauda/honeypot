FROM golang:1.14 as builder

WORKDIR /workspace
ADD . /workspace

ENV CGO_ENABLED=0

RUN go build -o honeypot .

FROM alpine:3.10

EXPOSE 8080
ENTRYPOINT ["/honeypot"]
RUN mkdir -p /logs

COPY --from=builder /workspace/honeypot /honeypot
