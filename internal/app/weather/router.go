package weather

import (
	"mweather-go/internal/app/weather/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(c *gin.Engine) {
	c.GET("/getweather", handler.SayWeather)
	c.GET("/weatherSearch", handler.SayCity)
}
