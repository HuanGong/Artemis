package avatar_data

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type (
	AvatarWeather struct {
		ErrorCode int    `json:"error_code"`
		Reason    string `json:"reason"`
		Result    struct {
			IsForeign int `json:"isForeign"`
			Life      struct {
				Date string `json:"date"`
				Info struct {
					Chuanyi   []string `json:"chuanyi"`
					Ganmao    []string `json:"ganmao"`
					Kongtiao  []string `json:"kongtiao"`
					Xiche     []string `json:"xiche"`
					Yundong   []string `json:"yundong"`
					Ziwaixian []string `json:"ziwaixian"`
				} `json:"info"`
			} `json:"life"`
			Pm25 struct {
				CityName string `json:"cityName"`
				DateTime string `json:"dateTime"`
				Key      string `json:"key"`
				Pm25     struct {
					CurPm   string `json:"curPm"`
					Des     string `json:"des"`
					Level   string `json:"level"`
					Pm10    string `json:"pm10"`
					Pm25    string `json:"pm25"`
					Quality string `json:"quality"`
				} `json:"pm25"`
			} `json:"pm25"`
			Realtime struct {
				CityCode   string `json:"city_code"`
				CityName   string `json:"city_name"`
				DataUptime string `json:"dataUptime"`
				Date       string `json:"date"`
				Moon       string `json:"moon"`
				Time       string `json:"time"`
				Weather    struct {
					Humidity    string `json:"humidity"`
					Img         string `json:"img"`
					Info        string `json:"info"`
					Temperature string `json:"temperature"`
				} `json:"weather"`
				Week string `json:"week"`
				Wind struct {
					Direct    string `json:"direct"`
					Offset    string `json:"offset"`
					Power     string `json:"power"`
					Windspeed string `json:"windspeed"`
				} `json:"wind"`
			} `json:"realtime"`
			Weather []struct {
				Date string `json:"date"`
				Info struct {
					Dawn  []string `json:"dawn"`
					Day   []string `json:"day"`
					Night []string `json:"night"`
				} `json:"info"`
				Nongli string `json:"nongli"`
				Week   string `json:"week"`
			} `json:"weather"`
		} `json:"result"`
	}
)

func GetWeather(city string, apikey string) (*AvatarWeather, error) {
	weather := &AvatarWeather{}
	fmtApi := "http://api.avatardata.cn/Weather/Query?key=%s&cityname=%s"

	fullApi := fmt.Sprintf(fmtApi, apikey, city)
	resp, err := http.Get(fullApi)
	if err != nil {
		return weather, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return weather, errors.Wrapf(err, "读取响应失败")
	}

	if err := json.Unmarshal(body, weather); err != nil {
		return weather, errors.Wrapf(err, "响应格式错误")
	}

	return weather, nil
}
