package app

import (
	"mweather-go/internal/app/healthz"
	"mweather-go/internal/app/weather"

	"github.com/gin-gonic/gin"
)

func registerRouters(router *gin.Engine) {
	healthz.RegisterRouters(router)
	weather.RegisterRouter(router)
}
