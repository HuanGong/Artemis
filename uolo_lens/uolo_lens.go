package uolo_lens

import (
	"github.com/sirupsen/logrus"
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/BurntSushi/toml"
	"encoding/json"
	"github.com/pkg/errors"
)

var (
	conf Config
	orm  *xorm.Engine
	WhiteList    map[string]bool
)

type (

	UoloLens struct {
		postHandler  *PostHandler
	}
)

func NewUoloLens() *UoloLens {
	WhiteList =    make(map[string]bool)

	instance := &UoloLens{
		postHandler: &PostHandler{},
	}

	loadConfig()

	initMysqlDB()

	return instance
}

func (impl *UoloLens) BeforeCliRun() error {
	return nil
}

func (impl *UoloLens) OnCliApplicationRun() error {
	return nil
}

func (impl *UoloLens) Endpoint() string {
	return conf.ServerAddress
}
func (impl *UoloLens) HttpServerName() string {
	return "UoloLens"
}

func (impl *UoloLens) OnServerInitialized(ec *echo.Echo) error {


	ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.HEAD, echo.DELETE, echo.OPTIONS},
	}))
/*
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
*/


	memGr := ec.Group("/lens")
	memGr.GET("/post/content", impl.postHandler.GetPostContent)

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

func initMysqlDB() error {

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
		return errors.Wrapf(err, "DB connection initialization failed")
	}

	orm = db

	if err := orm.Ping(); err != nil {
		logrus.Errorln("Ping Mysql Failed, Please Check Your Connection Config")
		return err
	}

	return nil
}


