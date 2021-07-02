package model

type City struct {
	Woeid         int     `json:"woeid"`         // 城市id
	Lat           float64 `json:"lat"`           // 经度
	Lon           float64 `json:"lon"`           // 纬度
	Country       string  `json:"country"`       // 国家
	City          string  `json:"city"`          // 城市名字
	QualifiedName string  `json:"qualifiedName"` // 官方城市名字
}
