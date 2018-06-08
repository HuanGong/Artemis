package uolo_lens

import (
	"artemis/uolo_lens/recommendation"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/importcjj/trie-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	LensConfig     Config
	Orm            *xorm.Engine
	JwtBlockTrie   *trie.Trie               = trie.New()
	RedisClientMap map[string]*redis.Client = make(map[string]*redis.Client)
	Recommender    *recommendation.RecommendImpl
)

type (
	UoloLens struct {
		postHandler   *PostHandler
		lensHandler   *LensHandler
		toolsHandler  *ToolsHandler
		thingsHandler *ThingsHandler
	}
)

func NewUoloLens() *UoloLens {

	instance := &UoloLens{
		postHandler:   &PostHandler{},
		lensHandler:   &LensHandler{},
		toolsHandler:  &ToolsHandler{},
		thingsHandler: &ThingsHandler{},
	}

	loadConfig()

	initMysqlDB()

	initRedis(LensConfig.RedisConfig)

	Recommender = recommendation.NewRecommendImpl(instance)

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
		Skipper:    impl.JwtSkipperChecker,
	}))
	ec.HTTPErrorHandler = func(err error, ec echo.Context) {
		logrus.Debugln(fmt.Sprintf("Request %s from ip %s occur error %s", ec.Path(), ec.RealIP(), err.Error()))
	}

	ec.GET("/lenslist", impl.lensHandler.LensList) // public

	utilsGr := ec.Group("/tools") // public
	utilsGr.GET("/weather/query", impl.toolsHandler.QueryWeather)
	utilsGr.GET("/extract/article", impl.postHandler.DialysisConent)

	ArticleGr := ec.Group("/article")
	ArticleGr.POST("/new", impl.postHandler.ArticleNew)
	ArticleGr.POST("/mod", impl.postHandler.ArticleMod)
	ArticleGr.GET("/detail", impl.postHandler.ArticleDetail)
	ArticleGr.GET("/auto/publish", impl.postHandler.AutoPublish)

	// JWT Auth Needed Path
	JwtBlockTrie.Put("/article/new", true)
	JwtBlockTrie.Put("/article/mod", true)
	JwtBlockTrie.Put("/article/auto/publish", true)

	ThingsGr := ec.Group("/things/v1")
	ThingsGr.GET("/list/todo", impl.thingsHandler.GetThingsList)         //public
	ThingsGr.GET("/list/finished", impl.thingsHandler.GetDoneThingsList) //public

	ThingsGr.POST("/new", impl.thingsHandler.NewThings)
	ThingsGr.POST("/delete", impl.thingsHandler.DeleteUoloThing)
	ThingsGr.POST("/finish", impl.thingsHandler.MarkUoloThingFinish)
	JwtBlockTrie.Put("/things/v1/new", true)
	JwtBlockTrie.Put("/things/v1/finish", true)
	JwtBlockTrie.Put("/things/v1/delete", true)

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

func (impl *UoloLens) DBOrmEngine() *xorm.Engine {
	return Orm
}

func (impl *UoloLens) RedisClient(name string) *redis.Client {
	c, _ := RedisClientMap[name]
	return c
}

func (impl *UoloLens) JwtSkipperChecker(ec echo.Context) bool {
	path := ec.Path()

	if ok, result := JwtBlockTrie.Match(path); ok && result.Value.(bool) == true {
		return false
	}

	return false
}
