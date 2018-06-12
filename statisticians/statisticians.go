package statisticians

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"strings"
)

var Conf Config = Config{}
var redisClient *redis.Client

type (
	Statistician struct {
		handler *StatisticHandler
	}
)

func NewStatistician() *Statistician {

	loadConfig()
	redisOption := &redis.Options{
		Addr:       "127.0.0.1:6379",
		DB:         1,
		MaxRetries: 0,
	}
	redisClient = redis.NewClient(redisOption)
	if _, err := redisClient.Ping().Result(); err != nil {
		logrus.Panicln("Connection Redis Failed")
		return nil
	}

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

func (impl *Statistician) OnHttpServerInitialized(ec *echo.Echo) error {

	ec.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(Conf.JWTSecretkey),
		Claims:     jwt.MapClaims{},
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
