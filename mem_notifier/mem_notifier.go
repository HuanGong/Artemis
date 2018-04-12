package mem_notifier

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"fmt"
	"github.com/labstack/echo/middleware"
)

type (

	Config struct {
		serverAddress string
	}

	Notifier struct {
		conf	*Config
		handler *Handler
	}
)

func NewNotifier() *Notifier {
	return &Notifier {
		handler: &Handler{},
	}
}

func (notifier *Notifier) BeforeCliRun() error {
	return nil
}

func (notifier *Notifier) OnCliApplicationRun() error {
	return nil
}

func (notifier *Notifier) Endpoint() string {
	return "0.0.0.0:3002"
}
func (notifier *Notifier) HttpServerName() string {
	return "MemNotifier"
}

func (notifier *Notifier) OnServerInitialized(ec *echo.Echo) error {

	ec.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	ec.GET("/", HandleWebHome)
	ec.POST("/signup", notifier.handler.SignUp)
	//ec.GET("/signup", notifier.handler.SignUp)
	ec.POST("/login", notifier.handler.Login)
	return nil
}

func HandleWebHome(ctx echo.Context) error {
	bytes, err := ioutil.ReadFile("/home/gh/www/index.html")


	fmt.Println(string(bytes))
	if err != nil {
		fmt.Println("return not find")
		ctx.String(404,"")
		return nil
	}
	fmt.Println("return html file")
	ctx.HTML(200, string(bytes))
	return nil
}

