package uolo_center

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

var (
	conf Config
	orm  *xorm.Engine
)

type (
	UoloCenter struct {
		handler      *Handler
		utilsHandler *UtilsHandler
		WhiteList    map[string]bool
	}
)

func NewNotifier() *UoloCenter {

	instance := &UoloCenter{
		handler:      &Handler{},
		utilsHandler: NewUtilsHandler(),
		WhiteList:    make(map[string]bool),
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
	return "UoloCenter"
}

func (impl *UoloCenter) OnServerInitialized(ec *echo.Echo) error {


	ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.HEAD, echo.DELETE, echo.OPTIONS},
	}))

	ec.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf.JWTSecretkey),
		Claims:     jwt.MapClaims{},
		Skipper: func(c echo.Context) bool {
			path := c.Path()

			logrus.Debugln("Request Path:", c.Path())

			if strings.HasPrefix(path, "/utils") {
				logrus.Debugln("skipper jwt verify")
				return true
			} else if strings.HasPrefix(path, "/au/login") ||
				strings.HasPrefix(path, "/au/signup") {
				return true
			}
			return false
		},
	}))


	gr := ec.Group("/au") // Authorization relative
	gr.POST("/signup", impl.handler.SignUp)
	gr.POST("/login", impl.handler.Login)
	gr.POST("/check", impl.handler.AuthTest)
	gr.POST("/reauth", impl.handler.AuthRefresh)
	gr.POST("/reset/passwd", impl.handler.ResetPassword)

	gr.GET("/signup", impl.handler.SignUp)
	gr.GET("/login", impl.handler.Login)
	gr.GET("/check", impl.handler.AuthTest)
	gr.GET("/reauth", impl.handler.AuthRefresh)
	gr.GET("/reset/passwd", impl.handler.ResetPassword)

	ec.GET("/cm", func(ec echo.Context) error {
		uidCookie := &http.Cookie{
			Name:     "UoloAU",
			Value:    "HuanGong",
			HttpOnly: false,
			Domain:   "",
			MaxAge:   int(time.Hour * 2),
		}
		ec.SetCookie(uidCookie)
		return ec.String(200, "")
	})

	utilsGr := ec.Group("/utils")
	utilsGr.GET("/qrcode2", impl.utilsHandler.GenQrcode)

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
