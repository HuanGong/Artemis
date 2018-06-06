package weather_cn

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	res, err := GetWeather("长沙")
	fmt.Println(res, err)
}

func TestGetAvatarWeatherWithCache(t *testing.T) {
	res, err := GetAvatarWeatherWithCache("北京")
	fmt.Println("res:", res, "err:", err)
}
