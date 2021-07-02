package model

type Weather struct {
	WoeId           int    // 城市ID
	Name            string // 城市名称
	Lower           string // 城市名小写
	date            int    // 日期时间戳
	MinTemperatureC int    // 摄氏-最低温度
	MaxTemperatureC int    // 摄氏-最高温度
	MinTemperatureF int    // 华氏-最低温度
	MaxTemperatureF int    // 华氏-最高温度
	Today           string // 天气 Cloudy
	Humidity        string // 湿度 89%
	SunriseTime     string // 日出时间 '06:25'
	SunriseH        int    // 日出小时 6
	SunriseI        int    // 日出分钟 25
	SunsetTime      string // 日落时间 '20:13'
	SunsetH         int    // 日落小时 20
	SunsetI         int    // 日落分钟 13
}
