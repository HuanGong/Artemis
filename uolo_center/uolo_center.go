package uolo_center

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"strings"
)

var (
	appConf      Config
	ormEngine    *xorm.Engine
	redisStorage map[string]*redis.Client
)

type (
	UoloCenter struct {
		authhandler    *AuthHandler
		profileHandler *ProfileHandler
		utilsHandler   *UtilsHandler
	}
)

func init() {
	loadConfig()
	initStorageDB()
	initRedisDb(appConf.RedisConfig)
}

func NewUoloCenter() *UoloCenter {

	conf := OptionConfig{
		Path: appConf.ProfilePhotoDir,
	}

	instance := &UoloCenter{
		authhandler:    &AuthHandler{},
		utilsHandler:   NewUtilsHandler(),
		profileHandler: NewProfileHandler(conf),
	}

	return instance
}

func (impl *UoloCenter) BeforeCliRun() error {
	return nil
}

func (impl *UoloCenter) OnCliApplicationRun() error {
	return nil
}

func (impl *UoloCenter) Endpoint() string {
	return appConf.ServerAddress
}
func (impl *UoloCenter) HttpServerName() string {
	return "UoloCenter"
}

func (impl *UoloCenter) OnHttpServerInitialized(ec *echo.Echo) error {

	ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"https://www.echoface.cn",
			"https://api.echoface.cn",
			"https://blog.echoface.cn",
			"http://localhost:4200",
			"http://localhost:3005",
		},
		AllowMethods:     []string{echo.GET, echo.POST, echo.HEAD, echo.DELETE, echo.OPTIONS, echo.PUT},
		AllowCredentials: true,
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

	ec.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(appConf.JWTSecretkey),
		Claims:      jwt.MapClaims{},
		TokenLookup: "cookie:_Authorization",
		Skipper: func(c echo.Context) bool {
			path := c.Path()
			if strings.HasPrefix(path, "/utils") {
				return true
			}
			if strings.HasPrefix(path, "/au/login") ||
				strings.HasPrefix(path, "/au/signup") {
				return true
			}
			return false
		},
	}))

	//ec.Use(utils.ServeCookie)

	impl.profileHandler.RegisterRouter(ec)

	gr := ec.Group("/au") // Authorization relative
	gr.POST("/signup", impl.authhandler.SignUp)
	gr.POST("/login", impl.authhandler.Login)
	gr.POST("/check", impl.authhandler.AuthTest)
	gr.POST("/renew/auth", impl.authhandler.AuthRefresh)
	gr.POST("/reset/passwd", impl.authhandler.ResetPassword)

	gr.GET("/signup", impl.authhandler.SignUp)
	gr.GET("/login", impl.authhandler.Login)
	gr.GET("/check", impl.authhandler.AuthTest)
	gr.GET("/renew/auth", impl.authhandler.AuthRefresh)
	gr.GET("/reset/passwd", impl.authhandler.ResetPassword)

	utilsGr := ec.Group("/utils")
	utilsGr.GET("/qrcode", impl.utilsHandler.GenQrcode)
	utilsGr.GET("/captcha", impl.utilsHandler.GenCaptcha)
	utilsGr.GET("/captcha/img", impl.utilsHandler.GenCaptchaImage)
	utilsGr.GET("/cmtest", func(ec echo.Context) error {
		cookies := ec.Cookies()
		return ec.JSON(200, cookies)
	})

	return nil
}

func loadConfig() {
	if _, err := toml.DecodeFile("config.toml", &appConf); err != nil {
		logrus.Panicln(err.Error())
	}
	jc, _ := json.MarshalIndent(&appConf, "", "  ")
	logrus.Infoln("load config", string(jc))
	level, err := logrus.ParseLevel(appConf.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetLevel(level)
}

func initStorageDB() {

	connectStr := &mysql.Config{
		User:   appConf.MysqlConfig.User,
		Passwd: appConf.MysqlConfig.Passwd,
		Net:    "tcp",
		Addr:   appConf.MysqlConfig.HostPort,
		DBName: appConf.MysqlConfig.DBName,
		Params: map[string]string{
			"charset": "utf8",
		},
	}

	db, err := xorm.NewEngine("mysql", connectStr.FormatDSN())
	if err != nil {
		logrus.Panic("DB connection initialization failed, ", err)
	}

	ormEngine = db

	if err := ormEngine.Ping(); err != nil {
		logrus.Errorln("Ping Mysql Failed, Please Check Your Connection Config")
	}
	return
}

func initRedisDb(redisConf []RedisConfig) {
	redisStorage = make(map[string]*redis.Client)

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
		redisStorage[conf.Name] = rClient
	}

}
