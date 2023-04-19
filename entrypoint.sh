#!/usr/bin/env ash

export GIN_MODE=release
export APP_URL=http://localhost
export APP_PORT=8080

# Connection Type: sqlite / mysql / postgres / dynamo
export DB_CONNECTION=sqlite
export DB_HOST=localhost
export DB_PORT=
export DB_DATABASE=go-bookstore.db
export DB_USERNAME=root
export DB_PASSWORD=

export JWT_AUTH_USERNAME=devopscorner
export JWT_AUTH_PASSWORD=DevOpsCorner2023
export JWT_SECRET=s3cr3t

export LOG_LEVEL=INFO

export AWS_REGION=us-west-2
export AWS_ACCESS_KEY=YOUR_AWS_ACCESS_KEY
export AWS_SECRET_KEY_ID=YOUR_AWS_SECRET_KEY_ID

export OPENSEARCH_ENDPOINT=https://opensearch.us-west-2.es.amazonaws.com
export OPENSEARCH_USERNAME=devopscorner
export OPENSEARCH_PASSWORD=DevOpsCorner2023

export PROMETHEUS_ENDPOINT=http://localhost:9090

export GRAFANA_ENDPOINT=http://localhost:3000
export GRAFANA_API_KEY=YOUR_GRAFANA_API_KEY

export OTEL_INSTRUMENTATION_METRIC_ENABLED=true
export OTEL_INSTRUMENTATION_TRACE_ENABLED=true
export OTEL_INSTRUMENTATION_LOG_ENABLED=true

# Trace Type: xray / jaeger
export OTEL_INSTRUMENTATION_TRACE_NAME=xray

export OTEL_SERVICE_NAME=bookstore-adot
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317
export OTEL_EXPORTER_OTLP_INSECURE=true
export OTEL_EXPORTER_OTLP_HEADERS=
export OTEL_RESOURCE_ATTRIBUTES=

export XRAY_VERSION=latest

export JAEGER_AGENT_PORT=6831
# Sampler Type: const / probabilistic / rateLimiting / remote
export JAEGER_SAMPLER_TYPE=const
export JAEGER_SAMPLER_PARAM=1
export JAEGER_SAMPLER_MANAGER_HOST_PORT=
export JAEGER_REPORTER_LOG_SPANS=true
# Interval in seconds (5*time.Second)
export JAEGER_REPORTER_BUFFER_FLUSH_INTERVAL="5*time.Second"
export JAEGER_REPORTER_MAX_QUEUE_SIZE=100
export JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT=
export JAEGER_REPORTER_COLLECTOR_ENDPOINT=http://localhost:14268/api/traces
export JAEGER_REPORTER_COLLECTOR_PASSWORD=DevOpsCorner2023
export JAEGER_TAGS=golang,otel,restful,api,bookstore

/go/goapp