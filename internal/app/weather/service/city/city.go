package city

import (
	"encoding/json"
	"errors"
	"fmt"
	"mweather-go/internal/app/weather/model"
	"mweather-go/internal/pkg/utlil/ip"
	"net/http"
	"time"

	"github.com/corpix/uarand"
	"github.com/go-resty/resty/v2"
)

const (
	BaseUrl       = "https://www.yahoo.com/"
	SearchCityUrl = "news/_td/api/resource/WeatherSearch;text="
)

type Client struct {
	*resty.Client
}

func (c *Client) SearchCity(city string) ([]model.City, error) {
	var citys []model.City
	url := fmt.Sprintf("%s%s%s", BaseUrl, SearchCityUrl, city)
	ip := ip.RandIP()
	response, err := client.R().
		SetHeader("User-Agent", uarand.GetRandom()).
		SetHeader("CLIENT-IP", ip).
		SetHeader("X-FORWARDED-FOR", ip).
		Get(url)
	if response.StatusCode() != http.StatusOK || err != nil {
		return citys, errors.New("fail")
	}
	if err := json.Unmarshal(response.Body(), &citys); err != nil {
		return citys, errors.New("fail")
	}
	return citys, nil
}

var client = new(Client)

func Register() {
	client.Client = resty.New()
	client.SetRetryCount(5)
	client.SetRetryWaitTime(5 * time.Second)
	client.SetDebug(true)
}

func GetClient() *Client {
	return client
}
