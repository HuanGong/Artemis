package uolo_lens

import (
	"artemis/uolo_lens/recommendation"
	"artemis/uolo_lens/utils"
	"encoding/json"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"strings"
	"time"
)

var (
	LensConfig     Config
	Orm            *xorm.Engine
	RedisClientMap = make(map[string]*redis.Client)
	Recommender    *recommendation.RecommendImpl
)

var (
	uniformDomain     = flag.Float64("uniform.domain", 0.0002, "The domain for the uniform distribution.")
	normDomain        = flag.Float64("normal.domain", 0.0002, "The domain for the normal distribution.")
	normMean          = flag.Float64("normal.mean", 0.00001, "The mean for the normal distribution.")
	oscillationPeriod = flag.Duration("oscillation-period", 10*time.Minute, "The duration of the rate oscillation period.")

	rpcDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "rpc_durations_seconds",
			Help:       "RPC latency distributions.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"service"},
	)

	rpcDurationsHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "rpc_durations_histogram_seconds",
		Help:    "RPC latency distributions.",
		Buckets: prometheus.LinearBuckets(*normMean-5**normDomain, .5**normDomain, 20),
	})
)

type (
	UoloLens struct {
		postHandler   *PostHandler
		lensHandler   *LensHandler
		toolsHandler  *ToolsHandler
		thingsHandler *ThingsHandler
	}
)

func init() {
	loadConfig()

	initStorageDB()

	initRedis(LensConfig.RedisConfig)
}

func NewUoloLens() *UoloLens {

	instance := &UoloLens{
		postHandler:   &PostHandler{},
		lensHandler:   &LensHandler{},
		toolsHandler:  &ToolsHandler{},
		thingsHandler: &ThingsHandler{},
	}

	Recommender = recommendation.NewRecommendImpl(instance)

	return instance
}

func (impl *UoloLens) BeforeCliRun() error {
	prometheus.MustRegister(rpcDurations)
	prometheus.MustRegister(rpcDurationsHistogram)
	return nil
}

func (impl *UoloLens) OnCliApplicationRun() error {
	flag.Parse()

	start := time.Now()
	oscillationFactor := func() float64 {
		return 2 + math.Sin(math.Sin(2*math.Pi*float64(time.Since(start))/float64(*oscillationPeriod)))
	}

	// Periodically record some sample latencies for the three services.
	go func() {
		for {
			v := rand.Float64() * *uniformDomain
			rpcDurations.WithLabelValues("uniform").Observe(v)
			time.Sleep(time.Duration(100*oscillationFactor()) * time.Millisecond)
		}
	}()

	go func() {
		for {
			v := (rand.NormFloat64() * *normDomain) + *normMean
			rpcDurations.WithLabelValues("normal").Observe(v)
			rpcDurationsHistogram.Observe(v)
			time.Sleep(time.Duration(75*oscillationFactor()) * time.Millisecond)
		}
	}()

	go func() {
		for {
			v := rand.ExpFloat64() / 1e6
			rpcDurations.WithLabelValues("exponential").Observe(v)
			time.Sleep(time.Duration(50*oscillationFactor()) * time.Millisecond)
		}
	}()

	return nil
}

func (impl *UoloLens) Endpoint() string {
	return LensConfig.ServerAddress
}
func (impl *UoloLens) HttpServerName() string {
	return "UoloLens"
}

func (impl *UoloLens) OnHttpServerInitialized(ec *echo.Echo) error {

	//ec.Use(middleware.CSRF())
	ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"https://www.echoface.cn",
			"https://blog.echoface.cn",
			"https://api.echoface.cn",
			"http://localhost:4200",
			"http://localhost:3005",
		},
		AllowMethods:     []string{echo.GET, echo.POST, echo.HEAD, echo.DELETE, echo.OPTIONS, echo.PUT},
		AllowCredentials: true,
	}))

	ec.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(LensConfig.JWTSecretkey),
		Claims:      jwt.MapClaims{},
		TokenLookup: "cookie:_Authorization",
		Skipper:     impl.JwtSkipperChecker,
	}))

	//ec.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	//	return func(ec echo.Context) error {
	//		logrus.Debugln("Cookie Viewer MiddleFunc >>>>>>>>>>>>>>>>>>>>>>> ")
	//		for _, cookie := range ec.Cookies() {
	//			logrus.Infof(" ===> %s : %s", cookie.Name, cookie.Value)
	//		}
	//		logrus.Debugln("Cookie Viewer MiddleFunc <<<<<<<<<<<<<<<<<<<<<<< ")
	//		return next(ec)
	//	}
	//})

	// public

	ec.GET("/p/lenslist", impl.lensHandler.LensList)
	ec.GET("/p/article/detail", impl.postHandler.ArticleDetail)
	ec.GET("/p/article/auto/publish", impl.postHandler.AutoPublish)
	ec.POST("/p/article/detail", impl.postHandler.ArticleDetail)
	ec.POST("p/article/auto/publish", impl.postHandler.AutoPublish)

	ec.GET("/p/tools/cmtest", func(ec echo.Context) error {
		cookies := ec.Cookies()
		return ec.JSON(200, cookies)
	})
	ec.GET("/p/tools/weather/query", impl.toolsHandler.QueryWeather)
	ec.GET("/p/tools/extract/article", impl.postHandler.DialysisConent)
	ec.GET("/p/things/v1/public/list", impl.thingsHandler.GetPublicThingsList)

	//private
	ec.GET("/lenslist", impl.lensHandler.LensList)
	ec.POST("/article/auto/publish", impl.postHandler.AutoPublish)

	ThingsGr := ec.Group("/things/v1")
	//ThingsGr.POST("/new", impl.thingsHandler.NewThings)
	//ThingsGr.POST("/delete", impl.thingsHandler.DeleteUoloThing)
	ThingsGr.POST("/finish", impl.thingsHandler.MarkUoloThingFinish)
	ThingsGr.GET("/list/todo", impl.thingsHandler.GetThingsList)
	ThingsGr.GET("/list/finished", impl.thingsHandler.GetDoneThingsList)

	return nil
}

func loadConfig() {
	if _, err := toml.DecodeFile("config.toml", &LensConfig); err != nil {
		logrus.Panicln(err.Error())
	}
	jc, _ := json.MarshalIndent(&LensConfig, "", "  ")
	logrus.Infoln("load config", string(jc))
	level, err := logrus.ParseLevel(LensConfig.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if LensConfig.CleanMarkPath != "" {
		utils.CleanMarkPath = LensConfig.CleanMarkPath
	}
	logrus.SetLevel(level)
}

func initStorageDB() error {

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

	if strings.HasPrefix(path, "/p/") {
		return true
	}

	return false
}
