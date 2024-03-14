ARG VARNISH_TAG=6.6
FROM golang:1.22.1-bookworm as builder
ARG CGO_ENABLED=0
WORKDIR /go/src/github.com/newrelic/nri-varnish
COPY . .
RUN make clean compile

# We use the varnish image as a base because the nri-varnish integration executes the 
# varnishstat command inside the varnish image to collect metrics.
FROM varnish:$VARNISH_TAG
COPY --from=builder /go/src/github.com/newrelic/nri-varnish/bin /
