# Golang Bookstore ADOT - Test RESTful API

ADOT (AWS Distro for OpenTelemetry) Implementation for Simple Golang RESTful API Application (Bookstore)

![goreport](https://goreportcard.com/badge/github.com/devopscorner/golang-adot/src)
![all contributors](https://img.shields.io/github/contributors/devopscorner/golang-adot)
![tags](https://img.shields.io/github/v/tag/devopscorner/golang-adot?sort=semver)
[![docker pulls](https://img.shields.io/docker/pulls/devopscorner/bookstore-adot.svg)](https://hub.docker.com/r/devopscorner/bookstore-adot/)
![download all](https://img.shields.io/github/downloads/devopscorner/golang-adot/total.svg)
![view](https://views.whatilearened.today/views/github/devopscorner/golang-adot.svg)
![clone](https://img.shields.io/badge/dynamic/json?color=success&label=clone&query=count&url=https://github.com/devopscorner/golang-adot/blob/master/clone.json?raw=True&logo=github)
![issues](https://img.shields.io/github/issues/devopscorner/golang-adot)
![pull requests](https://img.shields.io/github/issues-pr/devopscorner/golang-adot)
![forks](https://img.shields.io/github/forks/devopscorner/golang-adot)
![stars](https://img.shields.io/github/stars/devopscorner/golang-adot)
[![license](https://img.shields.io/github/license/devopscorner/golang-adot)](https://img.shields.io/github/license/devopscorner/golang-adot)

---

## Development

### Prequests

- Install jq libraries

  ```
  apt-get install -y jq
  ```

- Install golang dependencies

  ```
  cd src
  go mod init
  go mod tidy
  ```

### Runnning

```
go run main.go
```

### Runnning Test

```
go run main_test.go
```

## Folder Structure

```
.
├── .env
├── .env.example
├── config
│   ├── config.go
│   ├── config_test.go
│   ├── const.go
│   ├── logger.go
│   └── value.go
├── controller
│   ├── book_controller.go
│   ├── book_controller_test.go
│   ├── login_controller.go
│   └── login_controller_test.go
├── driver
│   ├── db.go
│   ├── dynamodb.go
│   ├── mysql.go
│   ├── psql.go
│   └── sqlite.go
├── go.mod
├── go.sum
├── main.go
├── main_test.go
├── middleware
│   ├── auth_middleware.go
│   └── auth_middleware_test.go
├── migrate_book.go.example
├── migrate_book_dynamo.go.example
├── model
│   └── book.go
├── observability
│   ├── metrics.go
│   ├── provider.go
│   ├── tracing-otel.go
│   └── tracing-xray.go
├── repository
│   └── book_repository.go
├── routes
│   ├── book_routes.go
│   ├── main_routes.go
│   ├── telemetry_routes.go
│   └── tracing_routes.go
└── view
    ├── book_view.go
    ├── error_view.go
    └── login_view.go

10 directories, 37 files
```

## Environment Variables (Default)

```
GIN_MODE=release
APP_URL=http://0.0.0.0
APP_PORT=8080

# Connection Type: sqlite / mysql / postgres / dynamo
DB_CONNECTION=dynamo
DB_HOST=0.0.0.0
DB_PORT=5000
DB_DATABASE=dynamodb-golang-adot
DB_USERNAME=root
DB_PASSWORD=

JWT_AUTH_USERNAME=devopscorner
JWT_AUTH_PASSWORD=DevOpsCorner2023
JWT_SECRET=s3cr3t

LOG_LEVEL=INFO

AWS_REGION=us-west-2
AWS_ACCESS_KEY=YOUR_AWS_KEY
AWS_SECRET_KEY_ID=YOUR_SECRET_KEY

OPENSEARCH_ENDPOINT=https://opensearch.us-west-2.es.amazonaws.com
OPENSEARCH_USERNAME=devopscorner
OPENSEARCH_PASSWORD=DevOpsCorner2023

PROMETHEUS_ENDPOINT=http://0.0.0.0:9090
PROMETHEUS_PORT=9090

GRAFANA_ENDPOINT=http://0.0.0.0:3000
GRAFANA_API_KEY=YOUR_GRAFANA_API_KEY

OTEL_INSTRUMENTATION_METRIC_ENABLED=true
OTEL_INSTRUMENTATION_TRACE_ENABLED=true
OTEL_INSTRUMENTATION_LOG_ENABLED=true

# Trace Type: xray / jaeger
OTEL_INSTRUMENTATION_TRACE_NAME=xray

OTEL_SERVICE_NAME=bookstore-adot
OTEL_EXPORTER_OTLP_ENDPOINT=http://0.0.0.0:4317
OTEL_EXPORTER_OTLP_PORT=4317
OTEL_EXPORTER_OTLP_INSECURE=true
OTEL_EXPORTER_OTLP_HEADERS=
OTEL_RESOURCE_ATTRIBUTES=
OTEL_TIME_INTERVAL=1
OTEL_RANDOM_TIME_ALIVE_INCREMENTER=1
OTEL_RANDOM_TOTAL_HEAP_SIZE_UPPER_BOUND=100
OTEL_RANDOM_THREAD_ACTIVE_UPPOR_BOUND=10
OTEL_RANDOM_CPU_USAGE_UPPER_BOUND=100

XRAY_VERSION=latest
XRAY_DAEMON_ENDPOINT=https://xray.us-west-2.amazonaws.com
XRAY_DAEMON_PORT=2000

JAEGER_AGENT_PORT=6831
# Sampler Type: const / probabilistic / rateLimiting / remote
JAEGER_SAMPLER_TYPE=const
JAEGER_SAMPLER_PARAM=1
JAEGER_SAMPLER_MANAGER_HOST_PORT=http://0.0.0.0:5778
JAEGER_REPORTER_LOG_SPANS=true
# Interval in seconds (5*time.Second)
JAEGER_REPORTER_BUFFER_FLUSH_INTERVAL="5*time.Second"
JAEGER_REPORTER_MAX_QUEUE_SIZE=100
JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT=http://0.0.0.0:6831
JAEGER_REPORTER_COLLECTOR_ENDPOINT=http://0.0.0.0:14268/api/traces
JAEGER_REPORTER_COLLECTOR_USER=devopscorner
JAEGER_REPORTER_COLLECTOR_PASSWORD=DevOpsCorner2023
JAEGER_TAGS=golang,otel,restful,api,bookstore
```

## Multi Driver Connection

```
DB_CONNECTION=sqlite
---
Available for:
- sqlite
- mysql
- postgres
- dynamo
```

## DynamoDB Connection

```
DB_CONNECTION=dynamo
---
DB_DATABASE --> Dynamo Table
DB_REGION   --> Dynamo Region
```

## API Test

- Generate JWT Token

```
POST    : /login
          curl --location '0.0.0.0:8080/v1/login' \
              --header 'Content-Type: application/json' \
              --data-raw '{
                  "username": "devopscorner",
                  "password": "DevOpsCorner2023"
              }' | jq
---
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzg2MTc5MzN9.p92DDMXVJPA8VTRDzDb-0NtzyfpdOtm5o6cJHMuZv44"
}

TOKEN=$(curl --request POST \
              --location '0.0.0.0:8080/v1/login' \
              --header 'Content-Type: application/json' \
              --data-raw '{
                  "username": "devopscorner",
                  "password": "DevOpsCorner2023"
              }' | jq -r '.token' )
---
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzg2MTc5MzN9.p92DDMXVJPA8VTRDzDb-0NtzyfpdOtm5o6cJHMuZv44
```

- List Books

```
GET     : /v1/books
          curl --request GET \
              --location '0.0.0.0:8080/v1/books' \
              --header 'Content-Type: application/json' \
              --header 'Authorization: Bearer ${TOKEN}'  | jq
```

- Find Book ID

```
GET     : /v1/books/${id}
          curl --request GET \
              --location '0.0.0.0:8080/v1/books/${id}' \
              --header 'Content-Type: application/json' \
              --header 'Authorization: Bearer ${TOKEN}'  | jq
```

- Add Book 1

```
POST    : /v1/books
          curl --request POST \
              --location '0.0.0.0:8080/v1/books' \
              --header 'Content-Type: application/json' \
              --header 'Authorization: Bearer ${TOKEN}' \
              --data-raw '{
                  "title": "Mastering Go: Create Golang production applications using network libraries, concurrency, and advanced Go data structures",
                  "author": "Mihalis Tsoukalos",
                  "year": "2023"
              }' | jq
```

- Add Book 2

```
POST    : /v1/books
          curl --request POST \
              --location '0.0.0.0:8080/v1/books' \
              --header 'Content-Type: application/json' \
              --header 'Authorization: Bearer ${TOKEN}' \
              --data-raw '{
                  "title": "Introducing Go: Build Reliable, Scalable Programs",
                  "author": "Caleb Doxsey",
                  "year": "2023"
              }' | jq
```

- Add Book 3

```
POST    : /v1/books
          curl --request POST \
              --location '0.0.0.0:8080/v1/books' \
              --header 'Content-Type: application/json' \
              --header 'Authorization: Bearer ${TOKEN}' \
              --data-raw '{
                  "title": "Learning Functional Programming in Go: Change the way you approach your applications using functional programming in Go",
                  "author": "Lex Sheehan",
                  "year": "2023"
              }' | jq
```

- Edit Book 3

```
PATCH   : /v1/books/3
          curl --request PATCH \
              --location '0.0.0.0:8080/v1/books/3' \
              --header 'Content-Type: application/json' \
              --header 'Authorization: Bearer ${TOKEN}' \
              --data-raw '{
                  "title": "DevOpsCorner",
                  "author": "DevOpsCorner Indonesia",
                  "year": "2023"
              }' | jq
```

- Delete Book 3

```
DELETE  : /v1/books/3
          curl --request DELETE \
              --location '0.0.0.0:8080/v1/books/3' \
              --header 'Content-Type: application/json' \
              --header 'Authorization: Bearer ${TOKEN}' | jq
```
