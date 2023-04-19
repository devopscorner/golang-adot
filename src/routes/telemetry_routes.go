package routes

import (
	"log"
	"net/http"

	"github.com/devopscorner/golang-adot/src/config"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func TelemetryRoutes(router *gin.Engine) {
	SetOtelRoutes(router)
	SetMetricsRoutes(router)
}

func SetOtelRoutes(router *gin.Engine) {
	// The Router
	router.Use(otelgin.Middleware(config.OtelServiceName()))
}

func SetMetricsRoutes(router *gin.Engine) {
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	go func() {
		log.Println("Prometheus metrics running on :" + config.PrometheusPort())
		log.Fatal(http.ListenAndServe(":"+config.PrometheusPort(), nil))
	}()
}
