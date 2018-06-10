package uolo_lens

import (
	"artemis/component/weather_cn"
	"artemis/component/weather_cn/avatar_data"
	"artemis/uolo_lens/utils"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type (
	ToolsHandler struct {
	}
)

func (handler *ToolsHandler) QueryWeather(ec echo.Context) error {
	city := ec.QueryParam("city")
	if len(city) == 0 {
		city = utils.GetCityByIp(ec.RealIP())
	}

	if _, has := weather_cn.CityWeatherCodeMap[city]; !has {
		return ec.JSON(http.StatusOK, &avatar_data.AvatarWeather{
			ErrorCode: -1,
			Reason:    "获取地理位置失败",
		})
	}

	info, err := weather_cn.GetAvatarWeatherWithCache(city)
	if err != nil {
		return ec.JSON(http.StatusOK, &avatar_data.AvatarWeather{
			ErrorCode: -2,
			Reason:    fmt.Sprintf("获取%s天气失败", city),
		})
	}
	return ec.JSON(http.StatusOK, info)
}
