FROM golang:1.13-alpine as builder

WORKDIR /go/src/app
COPY main.go .
COPY go.mod .
COPY go.sum .

RUN GCO_ENABLED=0 go install --ldflags '-extldflags "-static"'

# TODO: fix scratch!
FROM golang:1.13-alpine
COPY --from=builder /go/bin/knative-receiver /knative-receiver
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/knative-receiver"]