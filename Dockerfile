FROM golang:1.18.1 as builder

RUN mkdir -p /api
ADD . /api
WORKDIR /api

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o service ./cmd

FROM scratch

COPY --from=builder /api /api
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

CMD ["/api/service"]
