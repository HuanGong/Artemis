package statisticians

import (
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
	"strings"
)

type (
	Statistician struct {
		handler *StatisticHandler
	}
)

func NewStatistician() *Statistician {
	return &Statistician{
		handler: &StatisticHandler{},
	}
}

func (impl *Statistician) BeforeCliRun() error {
	return nil
}

func (impl *Statistician) OnCliApplicationRun() error {
	return nil
}

func (impl *Statistician) Endpoint() string {
	return "0.0.0.0:3005"
}
func (impl *Statistician) HttpServerName() string {
	return "Statistician"
}

func (impl *Statistician) OnServerInitialized(ec *echo.Echo) error {

	ec.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("123"),
		Claims:  jwt.MapClaims{},
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Path(), "/pri")
		},
	}))

	ec.GET("/statistics/pv", impl.handler.PVStatistic)

	return nil
}


