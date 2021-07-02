package app

import (
	"mweather-go/internal/app/healthz"

	"github.com/gin-gonic/gin"
)

func registerRouters(router *gin.Engine) {
	healthz.RegisterRouters(router)
}
