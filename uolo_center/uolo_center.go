package uolo_center

import (
	"github.com/labstack/echo"
	"github.com/go-xorm/xorm"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"github.com/labstack/echo/middleware"
)

var (
	conf Config
	orm  *xorm.Engine
)

type (

	UoloCenter struct {
		handler 	*Handler
		passPath 	map[string]bool
	}
)

func NewNotifier() *UoloCenter {

	instance :=  &UoloCenter{
		handler: 	&Handler{},
		passPath: 	map[string]bool {"/au/login":true, "/au/signup": true},
	}

	loadConfig()

	initDB()

	return instance
}

func (impl *UoloCenter) BeforeCliRun() error {
	return nil
}

func (impl *UoloCenter) OnCliApplicationRun() error {
	return nil
}

func (impl *UoloCenter) Endpoint() string {
	return conf.ServerAddress
}
func (impl *UoloCenter) HttpServerName() string {
	return "MemReminder"
}

func (impl *UoloCenter) OnServerInitialized(ec *echo.Echo) error {

	ec.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf.JWTSecretkey),
		Claims:  jwt.MapClaims{},
		Skipper: func(c echo.Context) bool {
			_, ok := impl.passPath[c.Path()]
			return ok
		},
	}))

	gr := ec.Group("/au") // Authorization relative
	gr.POST("/signup", impl.handler.SignUp)
	gr.POST("/login", impl.handler.Login)
	gr.POST("/check", impl.handler.AuthTest)
	gr.POST("/reauth", impl.handler.AuthRefresh)

	return nil
}

func loadConfig() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		logrus.Panicln(err.Error())
	}
	jc, _ := json.Marshal(&conf)
	logrus.Infoln("load config.toml", string(jc))
	level, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetLevel(level)
}

func initDB() {

	connectStr := &mysql.Config{
		User:   conf.MysqlConfig.User,
		Passwd: conf.MysqlConfig.Passwd,
		Net:    "tcp",
		Addr:   conf.MysqlConfig.HostPort,
		DBName: conf.MysqlConfig.DBName,
		Params: map[string]string{
			"charset": "utf8",
		},
	}

	db, err := xorm.NewEngine("mysql", connectStr.FormatDSN())
	if err != nil {
		logrus.Panic("DB connection initialization failed, ", err)
	}

	orm = db

	if err := orm.Ping(); err != nil {
		logrus.Errorln("Ping Mysql Failed, Please Check Your Connection Config")
	}
	return
}