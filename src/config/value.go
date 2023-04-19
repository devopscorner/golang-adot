package config

import (
	"github.com/spf13/viper"
)

// ----------------------------------------
// APPLICATION
// ----------------------------------------
func AppUrl() string {
	config := &Config{
		AppUrl: viper.GetString("APP_URL"),
	}
	return config.AppUrl
}

func AppPort() string {
	config := &Config{
		AppPort: viper.GetString("APP_PORT"),
	}
	return config.AppPort
}

// ----------------------------------------
// DATABASE
// ----------------------------------------
func DbConnection() string {
	config := &Config{
		DbConnection: viper.GetString("DB_CONNECTION"),
	}
	return config.DbConnection
}

func DbHost() string {
	config := &Config{
		DbHost: viper.GetString("DB_HOST"),
	}
	return config.DbHost
}

func DbPort() string {
	config := &Config{
		DbPort: viper.GetString("DB_PORT"),
	}
	return config.DbPort
}

func DbDatabase() string {
	config := &Config{
		DbDatabase: viper.GetString("DB_DATABASE"),
	}
	return config.DbDatabase
}

func DbUsername() string {
	config := &Config{
		DbUsername: viper.GetString("DB_USERNAME"),
	}
	return config.DbUsername
}

func DbPassword() string {
	config := &Config{
		DbPassword: viper.GetString("DB_PASSWORD"),
	}
	return config.DbPassword
}

// ----------------------------------------
// JWT (JSON Web Token)
// ----------------------------------------
func JWTAuthUsername() string {
	config := &Config{
		JwtAuthUsername: viper.GetString("JWT_AUTH_USERNAME"),
	}
	return config.JwtAuthUsername
}

func JWTAuthPassword() string {
	config := &Config{
		JwtAuthPassword: viper.GetString("JWT_AUTH_PASSWORD"),
	}
	return config.JwtAuthPassword
}

func JWTIssuer() string {
	config := &Config{
		JwtIssuer: viper.GetString("JWT_AUTH_USERNAME"),
	}
	return config.JwtIssuer
}

func JWTSecret() string {
	config := &Config{
		JwtSecret: viper.GetString("JWT_SECRET"),
	}
	return config.JwtSecret
}

// ----------------------------------------
// LOG Level
// ----------------------------------------
func LogLevel() string {
	config := &Config{
		LogLevel: viper.GetString("LOG_LEVEL"),
	}
	return config.LogLevel
}

// ----------------------------------------
// AWS Credentials
// ----------------------------------------
func AWSRegion() string {
	config := &Config{
		AWSRegion: viper.GetString("AWS_REGION"),
	}
	return config.AWSRegion
}

func AWSAccessKey() string {
	config := &Config{
		AWSAccessKey: viper.GetString("AWS_ACCESS_KEY"),
	}
	return config.AWSAccessKey
}

func AWSSecretKey() string {
	config := &Config{
		AWSSecretKey: viper.GetString("AWS_SECRET_KEY_ID"),
	}
	return config.AWSSecretKey
}

// -----------------------------------
// OPENSEARCH
// -----------------------------------
func OpenSearchEndpoint() string {
	config := &Config{
		OpenSearchEndpoint: viper.GetString("OPENSEARCH_ENDPOINT"),
	}
	return config.OpenSearchEndpoint
}

func OpenSearchUsername() string {
	config := &Config{
		OpenSearchUsername: viper.GetString("OPENSEARCH_USERNAME"),
	}
	return config.OpenSearchUsername
}

func OpenSearchPassword() string {
	config := &Config{
		OpenSearchPassword: viper.GetString("OPENSEARCH_PASSWORD"),
	}
	return config.OpenSearchPassword
}

// -----------------------------------
// PROMETHEUS
// -----------------------------------
func PrometheusEndpoint() string {
	config := &Config{
		PrometheusEndpoint: viper.GetString("PROMETHEUS_ENDPOINT"),
	}
	return config.PrometheusEndpoint
}

func PrometheusPort() string {
	config := &Config{
		PrometheusPort: viper.GetString("PROMETHEUS_PORT"),
	}
	return config.PrometheusPort
}

// -----------------------------------
// GRAFANA
// -----------------------------------
func GrafanaEndpoint() string {
	config := &Config{
		GrafanaEndpoint: viper.GetString("GRAFANA_ENDPOINT"),
	}
	return config.GrafanaEndpoint
}

func GrafanaApiKey() string {
	config := &Config{
		GrafanaApiKey: viper.GetString("GRAFANA_API_KEY"),
	}
	return config.GrafanaApiKey
}

// -----------------------------------
// OTEL (OpenTelemetry)
// -----------------------------------
func OtelMetricEnable() string {
	config := &Config{
		OtelMetricEnable: viper.GetString("OTEL_INSTRUMENTATION_METRIC_ENABLED"),
	}
	return config.OtelMetricEnable
}

func OtelTraceEnable() string {
	config := &Config{
		OtelTraceEnable: viper.GetString("OTEL_INSTRUMENTATION_TRACE_ENABLED"),
	}
	return config.OtelTraceEnable
}

func OtelTraceName() string {
	config := &Config{
		OtelTraceName: viper.GetString("OTEL_INSTRUMENTATION_TRACE_NAME"),
	}
	return config.OtelTraceName
}

func OtelLogEnable() string {
	config := &Config{
		OtelLogEnable: viper.GetString("OTEL_INSTRUMENTATION_LOG_ENABLED"),
	}
	return config.OtelLogEnable
}

func OtelServiceName() string {
	config := &Config{
		OtelServiceName: viper.GetString("OTEL_SERVICE_NAME"),
	}
	return config.OtelServiceName
}

func OtelOtlpEndpoint() string {
	config := &Config{
		OtelOtlpEndpoint: viper.GetString("OTEL_EXPORTER_OTLP_ENDPOINT"),
	}
	return config.OtelOtlpEndpoint
}

func OtelOtlpPort() string {
	config := &Config{
		OtelOtlpPort: viper.GetString("OTEL_EXPORTER_OTLP_PORT"),
	}
	return config.OtelOtlpPort
}

func OtelOtlpInsecure() string {
	config := &Config{
		OtelOtlpInsecure: viper.GetString("OTEL_EXPORTER_OTLP_INSECURE"),
	}
	return config.OtelOtlpInsecure
}

func OtelOtlpHeader() string {
	config := &Config{
		OtelOtlpHeader: viper.GetString("OTEL_EXPORTER_OTLP_HEADERS"),
	}
	return config.OtelOtlpHeader
}

func OtelAttributes() string {
	config := &Config{
		OtelAttributes: viper.GetString("OTEL_RESOURCE_ATTRIBUTES"),
	}
	return config.OtelAttributes
}

// -----------------------------------
// JAEGER
// -----------------------------------
func JaegerAgentPort() string {
	config := &Config{
		JaegerAgentPort: viper.GetString("JAEGER_AGENT_PORT"),
	}
	return config.JaegerAgentPort
}

func JaegerSamplerType() string {
	config := &Config{
		JaegerSamplerType: viper.GetString("JAEGER_SAMPLER_TYPE"),
	}
	return config.JaegerSamplerType
}

func JaegerSamplerParam() string {
	config := &Config{
		JaegerSamplerParam: viper.GetString("JAEGER_SAMPLER_PARAM"),
	}
	return config.JaegerSamplerParam
}

func JaegerSamplerManagerHostPort() string {
	config := &Config{
		JaegerSamplerManagerHostPort: viper.GetString("JAEGER_SAMPLER_MANAGER_HOST_PORT"),
	}
	return config.JaegerSamplerManagerHostPort
}

func JaegerReporterLogSpan() string {
	config := &Config{
		JaegerReporterLogSpan: viper.GetString("JAEGER_REPORTER_LOG_SPANS"),
	}
	return config.JaegerReporterLogSpan
}

func JaegerReporterBufferFlushInterval() string {
	config := &Config{
		JaegerReporterBufferFlushInterval: viper.GetString("JAEGER_REPORTER_BUFFER_FLUSH_INTERVAL"),
	}
	return config.JaegerReporterBufferFlushInterval
}

func JaegerReporterMaxQueueSize() string {
	config := &Config{
		JaegerReporterMaxQueueSize: viper.GetString("JAEGER_REPORTER_MAX_QUEUE_SIZE"),
	}
	return config.JaegerReporterMaxQueueSize
}

func JaegerReporterLocalAgentHostPort() string {
	config := &Config{
		JaegerReporterLocalAgentHostPort: viper.GetString("JAEGER_REPORTER_LOCAL_AGENT_HOST_PORT"),
	}
	return config.JaegerReporterLocalAgentHostPort
}

func JaegerReporterCollectorEndpoint() string {
	config := &Config{
		JaegerReporterCollectorEndpoint: viper.GetString("JAEGER_REPORTER_COLLECTOR_ENDPOINT"),
	}
	return config.JaegerReporterCollectorEndpoint
}

func JaegerReporterCollectorUser() string {
	config := &Config{
		JaegerReporterCollectorUser: viper.GetString("JAEGER_REPORTER_COLLECTOR_USER"),
	}
	return config.JaegerReporterCollectorUser
}

func JaegerReporterCollectorPassword() string {
	config := &Config{
		JaegerReporterCollectorPassword: viper.GetString("JAEGER_REPORTER_COLLECTOR_PASSWORD"),
	}
	return config.JaegerReporterCollectorPassword
}

func JaegerTags() string {
	config := &Config{
		JaegerTags: viper.GetString("JAEGER_TAGS"),
	}
	return config.JaegerTags
}

// -----------------------------------
// X-RAY
// -----------------------------------
func XRayVersion() string {
	config := &Config{
		XRayVersion: viper.GetString("XRAY_VERSION"),
	}
	return config.XRayVersion
}

func XRayDaemonEndpoint() string {
	config := &Config{
		XRayDaemonEndpoint: viper.GetString("XRAY_DAEMON_ENDPOINT"),
	}
	return config.XRayDaemonEndpoint
}

func XRayDaemonPort() string {
	config := &Config{
		XRayDaemonPort: viper.GetString("XRAY_DAEMON_PORT"),
	}
	return config.XRayDaemonPort
}
