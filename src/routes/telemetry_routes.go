package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/devopscorner/golang-adot/src/config"
	"github.com/devopscorner/golang-adot/src/observability"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func TelemetryRoutes(router *gin.Engine) {
	SetMetricsRoutes(router)
	observability.InitObservability(router)
	go func() {
		log.Println("OTEL Servicename: " + config.OtelServiceName())
		log.Println("XRay running on endpoint: " + config.XRayDaemonEndpoint())
		log.Println("XRay verion: " + config.XRayVersion())
		log.Println("Prometheus metrics running on :" + strconv.Itoa(config.PrometheusPort()))
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.PrometheusPort()), nil))
	}()
}

func SetMetricsRoutes(router *gin.Engine) {
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
