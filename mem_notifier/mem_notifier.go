package mem_notifier

import "github.com/labstack/echo"

type (

	Config struct {
		serverAddress string
	}

	Notifier struct {
		conf	*Config
	}
)


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

func (notifier *Notifier) OnServerInitialized(echo *echo.Echo) error {

	return nil
}

