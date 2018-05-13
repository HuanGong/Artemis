package uolo_lens

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"
)

var (
	LensConfig     Config
	Orm            *xorm.Engine
	WhiteList      map[string]bool          = make(map[string]bool)
	RedisClientMap map[string]*redis.Client = make(map[string]*redis.Client)
)

type (
	UoloLens struct {
		postHandler *PostHandler
		lensHandler *LensHandler
	}
)

func NewUoloLens() *UoloLens {

	instance := &UoloLens{
		postHandler: &PostHandler{},
		lensHandler: &LensHandler{},
	}

	loadConfig()

	initMysqlDB()

	initRedis(LensConfig.RedisConfig)

	return instance
}

func (impl *UoloLens) BeforeCliRun() error {
	return nil
}

func (impl *UoloLens) OnCliApplicationRun() error {
	return nil
}

func (impl *UoloLens) Endpoint() string {
	return LensConfig.ServerAddress
}
func (impl *UoloLens) HttpServerName() string {
	return "UoloLens"
}

func (impl *UoloLens) OnServerInitialized(ec *echo.Echo) error {

	ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.HEAD, echo.DELETE, echo.OPTIONS},
	}))

	ec.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(LensConfig.JWTSecretkey),
		Claims:     jwt.MapClaims{},
		Skipper: func(c echo.Context) bool {
			path := c.Path()

			logrus.Debugln("Request Path:", c.Path())
			return true

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

	ec.GET("/lenslist", impl.lensHandler.LensList)

	utilsGr := ec.Group("/utils")
	utilsGr.GET("/extract/article", impl.postHandler.DialysisConent)

	ArticleGr := ec.Group("/article")
	ArticleGr.POST("/new", impl.postHandler.ArticleNew)
	ArticleGr.POST("/mod", impl.postHandler.ArticleMod)
	ArticleGr.POST("/detail", impl.postHandler.ArticleDetail)
	ArticleGr.GET("/detail", impl.postHandler.ArticleDetail)
	ArticleGr.GET("/auto/publish", impl.postHandler.AutoPublish)

	return nil
}

func loadConfig() {
	if _, err := toml.DecodeFile("config.toml", &LensConfig); err != nil {
		logrus.Panicln(err.Error())
	}
	jc, _ := json.Marshal(&LensConfig)
	logrus.Infoln("load config.toml", string(jc))
	level, err := logrus.ParseLevel(LensConfig.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetLevel(level)
}

func initMysqlDB() error {

	connectStr := &mysql.Config{
		User:   LensConfig.MysqlConfig.User,
		Passwd: LensConfig.MysqlConfig.Passwd,
		Net:    "tcp",
		Addr:   LensConfig.MysqlConfig.HostPort,
		DBName: LensConfig.MysqlConfig.DBName,
		Params: map[string]string{
			"charset": "utf8",
		},
	}

	db, err := xorm.NewEngine("mysql", connectStr.FormatDSN())
	if err != nil {
		logrus.Panic("DB connection initialization failed, ", err)
		return errors.Wrapf(err, "DB connection initialization failed")
	}

	Orm = db

	if err := Orm.Ping(); err != nil {
		logrus.Errorln("Ping Mysql Failed, Please Check Your Connection Config")
		return err
	}

	return nil
}

func initRedis(redisConf []RedisConfig) {
	for _, conf := range redisConf {
		rClient := redis.NewClient(&redis.Options{
			Addr:     conf.Addr,
			DB:       int(conf.DbIndex),
			Password: conf.Password,
		})

		_, err := rClient.Ping().Result()
		if err != nil {
			logrus.Errorln("Connect To Redis Failed, server:", conf)
			continue
		}
		if len(conf.Name) == 0 {
			logrus.Errorln("Redis Config Must Need A Name")
			continue
		}
		RedisClientMap[conf.Name] = rClient
	}
}
