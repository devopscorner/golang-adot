package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppUrl                            string
	AppPort                           string
	DbConnection                      string
	DbHost                            string
	DbPort                            string
	DbDatabase                        string
	DbUsername                        string
	DbPassword                        string
	JwtAuthUsername                   string
	JwtAuthPassword                   string
	JwtIssuer                         string
	JwtSecret                         string
	LogLevel                          string
	AWSRegion                         string
	AWSAccessKey                      string
	AWSSecretKey                      string
	OpenSearchEndpoint                string
	OpenSearchUsername                string
	OpenSearchPassword                string
	PrometheusEndpoint                string
	PrometheusPort                    string
	GrafanaEndpoint                   string
	GrafanaApiKey                     string
	OtelMetricEnable                  string
	OtelTraceEnable                   string
	OtelTraceName                     string
	OtelLogEnable                     string
	OtelServiceName                   string
	OtelOtlpEndpoint                  string
	OtelOtlpPort                      string
	OtelOtlpInsecure                  string
	OtelOtlpHeader                    string
	OtelAttributes                    string
	JaegerAgentPort                   string
	JaegerSamplerType                 string
	JaegerSamplerParam                string
	JaegerSamplerManagerHostPort      string
	JaegerReporterLogSpan             string
	JaegerReporterBufferFlushInterval string
	JaegerReporterMaxQueueSize        string
	JaegerReporterLocalAgentHostPort  string
	JaegerReporterCollectorEndpoint   string
	JaegerReporterCollectorUser       string
	JaegerReporterCollectorPassword   string
	JaegerTags                        string
	XRayDaemonEndpoint                string
	XRayDaemonPort                    string
	XRayVersion                       string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.SetDefault("APP_URL", "http://localhost")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("DB_CONNECTION", "dynamo")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "")
	viper.SetDefault("DB_DATABASE", "dynamodb-golang-adot")
	viper.SetDefault("DB_USERNAME", "root")
	viper.SetDefault("DB_PASSWORD", "")
	viper.SetDefault("JWT_AUTH_USERNAME", "devopscorner")
	viper.SetDefault("JWT_AUTH_PASSWORD", "DevOpsCorner2023")
	viper.SetDefault("JWT_SECRET", "s3cr3t")

	// LOG_LEVEL: DEBUG | INFO | WARN | ERROR
	viper.SetDefault("LOG_LEVEL", "INFO")

	// AWS MANAGED SERVICES
	viper.SetDefault("AWS_REGION", "us-west-2")
	viper.SetDefault("AWS_ACCESS_KEY", "YOUR_AWS_ACCESS_KEY")
	viper.SetDefault("AWS_SECRET_KEY_ID", "YOUR_AWS_SECRET_KEY_ID")
	viper.SetDefault("OPENSEARCH_ENDPOINT", "https://opensearch.us-west-2.es.amazonaws.com")
	viper.SetDefault("OPENSEARCH_USERNAME", "OPENSEARCH_USERNAME")
	viper.SetDefault("OPENSEARCH_PASSWORD", "OPENSEARCH_PASSWORD")
	viper.SetDefault("PROMETHEUS_ENDPOINT", "http://localhost:9090")
	viper.SetDefault("PROMETHEUS_PORT", "9090")
	viper.SetDefault("GRAFANA_ENDPOINT", "http://localhost:3000")
	viper.SetDefault("GRAFANA_API_KEY", "GRAFANA_API_KEY")

	// OPEN TELEMETRY
	viper.SetDefault("OTEL_INSTRUMENTATION_METRIC_ENABLED", "true") // Prometheus Enable?
	viper.SetDefault("OTEL_INSTRUMENTATION_TRACE_ENABLED", "true")  // Tracing Enable?
	viper.SetDefault("OTEL_INSTRUMENTATION_LOG_ENABLED", "true")    // Logging Enable?
	viper.SetDefault("OTEL_SERVICE_NAME", "bookstore-adot")         // Service Name OTEL
	viper.SetDefault("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4317")
	viper.SetDefault("OTEL_EXPORTER_OTLP_PORT", "4317")
	viper.SetDefault("OTEL_EXPORTER_OTLP_INSECURE", "true")
	viper.SetDefault("OTEL_EXPORTER_OTLP_HEADERS", "")
	viper.SetDefault("OTEL_RESOURCE_ATTRIBUTES", "")

	// TRACING with XRAY
	viper.SetDefault("XRAY_VERSION", "latest")
	viper.SetDefault("XRAY_DAEMON_ENDPOINT", "http://localhost:4317")
	viper.SetDefault("XRAY_DAEMON_PORT", "2000")

	// TRACING with JAEGER
	viper.SetDefault("JAEGER_AGENT_PORT", "6831")
	viper.SetDefault("JAEGER_SAMPLER_TYPE", "const")
	viper.SetDefault("JAEGER_SAMPLER_PARAM", "1")
	viper.SetDefault("JAEGER_SAMPLER_MANAGER_HOST_PORT", "")
	viper.SetDefault("JAEGER_REPORTER_LOG_SPANS", "true")
	viper.SetDefault("JAEGER_REPORTER_BUFFER_FLUSH_INTERVAL", "5")
	viper.SetDefault("JAEGER_REPORTER_MAX_QUEUE_SIZE", "100")
	viper.SetDefault("JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT", "")
	viper.SetDefault("JAEGER_REPORTER_COLLECTOR_ENDPOINT", "http://localhost:14268/api/traces")
	viper.SetDefault("JAEGER_REPORTER_COLLECTOR_USER", "")
	viper.SetDefault("JAEGER_REPORTER_COLLECTOR_PASSWORD", "")
	viper.SetDefault("JAEGER_TAGS", "golang,otel,restful,api,bookstore")

	config := &Config{
		AppUrl:                            viper.GetString("APP_URL"),
		AppPort:                           viper.GetString("APP_PORT"),
		DbConnection:                      viper.GetString("DB_CONNECTION"),
		DbHost:                            viper.GetString("DB_HOST"),
		DbPort:                            viper.GetString("DB_PORT"),
		DbDatabase:                        viper.GetString("DB_DATABASE"),
		DbUsername:                        viper.GetString("DB_USERNAME"),
		DbPassword:                        viper.GetString("DB_PASSWORD"),
		JwtAuthUsername:                   viper.GetString("JWT_AUTH_USERNAME"),
		JwtAuthPassword:                   viper.GetString("JWT_AUTH_PASSWORD"),
		JwtSecret:                         viper.GetString("JWT_SECRET"),
		AWSRegion:                         viper.GetString("AWS_REGION"),
		AWSAccessKey:                      viper.GetString("AWS_ACCESS_KEY"),
		AWSSecretKey:                      viper.GetString("AWS_SECRET_KEY_ID"),
		OpenSearchEndpoint:                viper.GetString("OPENSEARCH_ENDPOINT"),
		OpenSearchUsername:                viper.GetString("OPENSEARCH_USERNAME"),
		OpenSearchPassword:                viper.GetString("OPENSEARCH_PASSWORD"),
		PrometheusEndpoint:                viper.GetString("PROMETHEUS_ENDPOINT"),
		PrometheusPort:                    viper.GetString("PROMETHEUS_PORT"),
		GrafanaEndpoint:                   viper.GetString("GRAFANA_ENDPOINT"),
		GrafanaApiKey:                     viper.GetString("GRAFANA_API_KEY"),
		OtelMetricEnable:                  viper.GetString("OTEL_INSTRUMENTATION_METRIC_ENABLED"),
		OtelTraceEnable:                   viper.GetString("OTEL_INSTRUMENTATION_TRACE_ENABLED"),
		OtelTraceName:                     viper.GetString("OTEL_INSTRUMENTATION_TRACE_NAME"),
		OtelLogEnable:                     viper.GetString("OTEL_INSTRUMENTATION_LOG_ENABLED"),
		OtelServiceName:                   viper.GetString("OTEL_SERVICE_NAME"),
		OtelOtlpEndpoint:                  viper.GetString("OTEL_EXPORTER_OTLP_ENDPOINT"),
		OtelOtlpPort:                      viper.GetString("OTEL_EXPORTER_OTLP_PORT"),
		OtelOtlpInsecure:                  viper.GetString("OTEL_EXPORTER_OTLP_INSECURE"),
		OtelOtlpHeader:                    viper.GetString("OTEL_EXPORTER_OTLP_HEADERS"),
		OtelAttributes:                    viper.GetString("OTEL_RESOURCE_ATTRIBUTES"),
		JaegerAgentPort:                   viper.GetString("JAEGER_AGENT_PORT"),
		JaegerSamplerType:                 viper.GetString("JAEGER_SAMPLER_TYPE"),
		JaegerSamplerParam:                viper.GetString("JAEGER_SAMPLER_PARAM"),
		JaegerSamplerManagerHostPort:      viper.GetString("JAEGER_SAMPLER_MANAGER_HOST_PORT"),
		JaegerReporterLogSpan:             viper.GetString("JAEGER_REPORTER_LOG_SPANS"),
		JaegerReporterBufferFlushInterval: viper.GetString("JAEGER_REPORTER_BUFFER_FLUSH_INTERVAL"),
		JaegerReporterMaxQueueSize:        viper.GetString("JAEGER_REPORTER_MAX_QUEUE_SIZE"),
		JaegerReporterLocalAgentHostPort:  viper.GetString("JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT"),
		JaegerReporterCollectorEndpoint:   viper.GetString("JAEGER_REPORTER_COLLECTOR_ENDPOINT"),
		JaegerReporterCollectorUser:       viper.GetString("JAEGER_REPORTER_COLLECTOR_USER"),
		JaegerReporterCollectorPassword:   viper.GetString("JAEGER_REPORTER_COLLECTOR_PASSWORD"),
		JaegerTags:                        viper.GetString("JAEGER_TAGS"),
		XRayVersion:                       viper.GetString("XRAY_VERSION"),
		XRayDaemonEndpoint:                viper.GetString("XRAY_DAEMON_ENDPOINT"),
		XRayDaemonPort:                    viper.GetString("XRAY_DAEMON_PORT"),
	}

	return config, nil
}
