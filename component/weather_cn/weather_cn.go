package weather_cn

import (
	"artemis/component/weather_cn/avatar_data"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
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

	AvatarWeatherCache struct {
		T       time.Time
		Weather *avatar_data.AvatarWeather
	}
)

const (
	apiPath = "http://www.weather.com.cn/data/sk/"
)

var (
	DebugEnable        bool              = false
	WeatherItems       []WeatherItem     = make([]WeatherItem, 0)
	CityWeatherCodeMap map[string]string = make(map[string]string)
	AvatarCacheMap     sync.Map
	AvatarDataApiKey   = "afdb952ba7aa4a419acd05e754587b17"
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

func GetAvatarWeatherWithCache(city string) (*avatar_data.AvatarWeather, error) {
	cache, has := AvatarCacheMap.Load(city)
	now := time.Now()
	if has {
		weatherCache := cache.(*AvatarWeatherCache)
		if weatherCache.T.Add(time.Hour * 4).After(now) {
			return weatherCache.Weather, nil
		}
	}

	weather, err := avatar_data.GetWeather(city, AvatarDataApiKey)
	if err != nil {
		return nil, err
	}

	NewCache := &AvatarWeatherCache{
		T:       now,
		Weather: weather,
	}
	AvatarCacheMap.Store(city, NewCache)
	return weather, nil
}
