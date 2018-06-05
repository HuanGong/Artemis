package weather_cn

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type (
	WeatherCnInfo struct {
		City     string `json:"city"`
		Temp     string `json:"temp"`
		Wind     string `json:"WD"`
		WdSpeed  string `json:"WS"`
		Humidity string `json:"SD"`
	}

	AvatarWeatherInfo struct {
		CityName      string
		Week          string
		Date          string
		RtWindInfo    string
		RtWindPower   string
		RtWeather     string
		RtTemperature string
		Pm25Level     string
		Pm25Current   string
		Pm25Quality   string
	}
)

const (
	apiPath = "http://www.weather.com.cn/data/sk/"
)

var (
	DebugEnable        bool              = false
	WeatherItems       []WeatherItem     = make([]WeatherItem, 0)
	CityWeatherCodeMap map[string]string = make(map[string]string)
)

func init() {
	err := json.Unmarshal([]byte(cityCodeCN), &WeatherItems)
	if err != nil {
		fmt.Println("Init Weather Meta Data Failed With Err ", err.Error())
	}

	for _, item := range WeatherItems {
		for _, city := range item.City {
			for _, county := range city.Counties {
				if _, ok := CityWeatherCodeMap[county.Name]; ok {
					fmt.Println("WeatherPackage Has Conflict on:", county.Name)
					continue
				}
				CityWeatherCodeMap[county.Name] = county.WeatherCode
			}
		}
	}
}

func GetWeather(city string) (WeatherCnInfo, error) {
	type (
		Res struct {
			Info WeatherCnInfo `json:"weatherinfo"`
		}
	)
	result := &Res{}

	code, ok := CityWeatherCodeMap[city]
	if !ok {
		code = "101010100"
	}
	res, err := http.Get(apiPath + code + ".html")
	if err != nil {
		return result.Info, errors.Wrap(err, "Get Weather info failed with network error")
	}
	defer res.Body.Close()
	body, Rerr := ioutil.ReadAll(res.Body)
	if Rerr != nil {
		return result.Info, errors.Wrap(err, "Read Weather Detail error")
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, result)
	return result.Info, err
}
