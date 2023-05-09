package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// --- X-Ray --- //
type XRayResponseWriter struct {
	gin.ResponseWriter
	parentName string
}

func (w XRayResponseWriter) Name(name string) string {
	// Use the response status code and parent segment name as the segment name
	return fmt.Sprintf("%s/%d", w.parentName, w.Status())
}
