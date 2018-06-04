package weather_cn

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	res, err := GetWeather("长沙")
	fmt.Println(res, err)
}
