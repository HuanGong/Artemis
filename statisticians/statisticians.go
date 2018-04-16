package statisticians

import (
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
	"strings"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"encoding/json"
)

var Conf Config = Config{}

type (
	Statistician struct {
		handler *StatisticHandler
	}
)

func NewStatistician() *Statistician {

	loadConfig()

	return &Statistician{
		handler: &StatisticHandler{},
	}
}

func loadConfig() {
	if _, err := toml.DecodeFile("config.toml", &Conf); err != nil {
		logrus.Panicln(err.Error())
	}
	jc, _ := json.Marshal(&Conf)
	logrus.Infoln("load config.toml", string(jc))
	level, err := logrus.ParseLevel(Conf.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetLevel(level)
}

func (impl *Statistician) BeforeCliRun() error {
	return nil
}

func (impl *Statistician) OnCliApplicationRun() error {
	return nil
}

func (impl *Statistician) Endpoint() string {
	return Conf.ServerAddress
}
func (impl *Statistician) HttpServerName() string {
	return "Statistician"
}

func (impl *Statistician) OnServerInitialized(ec *echo.Echo) error {

	ec.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(Conf.JWTSecretkey),
		Claims:  jwt.MapClaims{},
		Skipper: func(c echo.Context) bool {
			return !strings.HasPrefix(c.Path(), "/pri")
		},
	}))

	privateGr := ec.Group("/pri")
	privateGr.GET("/pv", impl.handler.PVStatistic)

	publicGr := ec.Group("/com")
	publicGr.GET("/pv", impl.handler.PVStatistic)

	return nil
}


