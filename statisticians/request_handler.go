package statisticians

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)
type (
	StatisticHandler struct {
	}
)

/* statistic the pv on site write down the the info ip,time as infomation*/
func (handler *StatisticHandler) PVStatistic(ctx echo.Context) error {

	logrus.Infoln("RealIP:", ctx.RealIP())
	logrus.Infoln("Path  :", ctx.Path())
	logrus.Infoln("Domain:", ctx.Request().Host)
	cookie, err := ctx.Cookie("uid")
	if err == nil {
		logrus.Infoln("UID   :", cookie.Value)
	}

	uidCookie := &http.Cookie{
		Name:   "uid",
		Value:    "gonghuan",
		HttpOnly: false,
		MaxAge:   int(time.Hour * 24),
	}

	ctx.SetCookie(uidCookie)
	ctx.String(200, "")

	return nil
}

