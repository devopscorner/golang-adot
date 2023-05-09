package observability

import (
	"github.com/gin-gonic/gin"
)

func InitObservability(router *gin.Engine) {
	InitMetrics(router)
	InitTracingXRay(router)
	InitTracingOtel(router)
}
