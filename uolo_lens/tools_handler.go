package uolo_lens

import (
	"artemis/component/weather_cn"
	"artemis/uolo_lens/utils"
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

	info, err := weather_cn.GetWeather(city)
	if err != nil {
		ec.JSON(http.StatusOK, &weather_cn.WeatherCnInfo{
			City:     "UOLO",
			Temp:     "FIT",
			Wind:     "清风",
			WdSpeed:  "微风",
			Humidity: "舒适",
		})
	}
	return ec.JSON(http.StatusOK, info)
}
