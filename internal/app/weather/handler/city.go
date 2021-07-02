package handler

import (
	"mweather-go/internal/app/weather/service/city"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/makeblock-go/log"
)

type City struct {
	WoeId int    `json:"woeid"` // 城市id
	Name  string `json:"name"`  // 城市名字
}

func SayCity(c *gin.Context) {
	cityName := c.Query("city")

	if IsCityNameEmpty(cityName) {
		c.JSON(http.StatusOK, []struct{}{})
		return
	}

	citys, err := city.GetClient().SearchCity(cityName)
	if err != nil {
		log.Print(err)
	}
	c.JSON(http.StatusOK, citys)
}

// IsCityNameEmpty 判断是否空
func IsCityNameEmpty(cityName string) bool {
	return cityName == ""
}
