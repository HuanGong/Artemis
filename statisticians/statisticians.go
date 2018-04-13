package statisticians

import "github.com/labstack/echo"


type (
	Statistician struct {
		handler *StatisticHandler
	}
)

func NewStatistician() *Statistician {
	return &Statistician{
		handler: &StatisticHandler{},
	}
}

func (impl *Statistician) BeforeCliRun() error {
	return nil
}

func (impl *Statistician) OnCliApplicationRun() error {
	return nil
}

func (impl *Statistician) Endpoint() string {
	return "0.0.0.0:3005"
}
func (impl *Statistician) HttpServerName() string {
	return "Statistician"
}

func (impl *Statistician) OnServerInitialized(ec *echo.Echo) error {

	ec.GET("/statistics/pv", impl.handler.PVStatistic)

	return nil
}


