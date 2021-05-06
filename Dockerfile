ARG GOLANG_VERSION=1.16

FROM golang:${GOLANG_VERSION} as builder
WORKDIR /code
COPY go.mod .
RUN go mod download

COPY . ./
RUN go build -o ./bin/nri-varnish cmd/nri-varnish/main.go; strip ./bin/nri-varnish

FROM newrelic/infrastructure:latest
ENV NRIA_IS_FORWARD_ONLY true
ENV NRIA_K8S_INTEGRATION true
COPY --from=builder /code/bin/nri-varnish /nri-sidecar/newrelic-infra/newrelic-integrations/bin/nri-postgresql
COPY --from=builder /code/varnish-definition.yml /nri-sidecar/newrelic-infra/newrelic-integrations/definition.yaml
USER 1000
