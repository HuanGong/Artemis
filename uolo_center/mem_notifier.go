package uolo_center

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/go-xorm/xorm"
	"github.com/go-sql-driver/mysql"
	"artemis/uolo_center/model"
	"github.com/sirupsen/logrus"
	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
)

var (
	conf Config
	orm  *xorm.Engine
)

type (

	Notifier struct {
		handler *Handler
	}
)

func NewNotifier() *Notifier {
	instance :=  &Notifier {
		handler: &Handler{},
	}
	loadConfig()

	initDB()

	return instance
}

func (notifier *Notifier) BeforeCliRun() error {
	return nil
}

func (notifier *Notifier) OnCliApplicationRun() error {
	return nil
}

func (notifier *Notifier) Endpoint() string {
	return conf.ServerAddress
}
func (notifier *Notifier) HttpServerName() string {
	return "MemReminder"
}

func (notifier *Notifier) OnServerInitialized(ec *echo.Echo) error {

	ec.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(JWTSecretkey),
		Claims:  jwt.MapClaims{},
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	ec.POST("/signup", notifier.handler.SignUp)
	ec.POST("/login", notifier.handler.Login)
	ec.POST("/re-auth", notifier.handler.AuthRefresh)
	ec.POST("/check-auth", notifier.handler.AuthTest)
	return nil
}

func loadConfig() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		logrus.Panicln(err.Error())
	}
	jc, _ := json.Marshal(&conf)
	logrus.Infoln("load config.toml", string(jc))
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

	//testId, _ := uuid.NewV4()

	_, err = orm.Insert(&model.User{
		Id: "TestID",
	})

	if err != nil {
		logrus.Errorln("Insert TestID to Mysql Failed， err：", err.Error())
	}
	return
}