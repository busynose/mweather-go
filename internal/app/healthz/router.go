package healthz

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouters register routers
func RegisterRouters(r *gin.Engine) {
	r.GET("/healthz", handleHealthz)
}
