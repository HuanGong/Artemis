package uolo

import (
	"os"
	"sync"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
	log "github.com/sirupsen/logrus"
)

type FlowRunType int32

const (
	kRunTypeSync  FlowRunType = 1
	kRunTypeAsync FlowRunType = 2
)

type (
	AppContent interface {
		BeforeCliRun() error
		OnCliApplicationRun() error
	}

	HttpServerContent interface {
		Endpoint() string
		HttpServerName() string
		OnServerInitialized(echo *echo.Echo) error
	}

	StartFlow struct {
		runType FlowRunType
		runFlow func() error
	}

	App struct {
		content     AppContent
		name        string
		version     string
		description string
		startFlow   []*StartFlow
		extraFlags  []cli.Flag
		subCommands []*cli.Command
	}
)

func NewApp(c AppContent) *App {
	app := &App{
		content: c,
	}
	return app
}

func (app *App) WithHttpServer(content HttpServerContent) *App {
	flow := &StartFlow{
		runType: kRunTypeAsync,
		runFlow: func() error {
			e := echo.New()
			e.HideBanner = true
			//e.Binder = &binder.PbRequestBinder{}

			//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
				//AllowOrigins: []string{"*"},
				//AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS, echo.DELETE, echo.HEAD},
			//}))

			if err := content.OnServerInitialized(e); err != nil {
				log.Errorf("App Run WithHttpServer %s Failed With Error %s\n", content.HttpServerName(), err.Error())
				return errors.Wrap(err, "Start HttpServer Failed")
			}
			return e.Start(content.Endpoint())
		},
	}

	app.startFlow = append(app.startFlow, flow)
	return app
}

func (app *App) WithCommonFlow(fun func() error, flowRunType FlowRunType) {
	flow := &StartFlow{
		runType: flowRunType,
		runFlow: fun,
	}
	app.startFlow = append(app.startFlow, flow)
}

func (app *App) CliAction(ctx *cli.Context) error {

	log.Infoln("Uolo Application Start Run CliAction")

	if app.content != nil {
		app.content.OnCliApplicationRun()
	}

	log.Infoln("start run start flow function")

	wg := sync.WaitGroup{}
	wg.Add(1) //avoid empty wg.wait
	for _, flow := range app.startFlow {
		if flow.runType == kRunTypeAsync {
			wg.Add(1)
			go func() {
				if err := flow.runFlow(); err != nil {
					log.Panicln("Start async Flow Failed", err.Error())
				}
				wg.Done()
			}()
		} else {
			wg.Add(1)
			if err := flow.runFlow(); err != nil {
				log.Panicln("Start sync Flow Failed", err.Error())
			}
			wg.Done()
		}
	}
	wg.Done()
	wg.Wait()

	log.Infoln("all start flow done!")
	return nil
}

func (app *App) Run() error {
	cli := &cli.App{
		Name:    app.name,
		Version: app.version,
		Usage:   app.description,
		Flags:   app.AddAppCliFlags(),
		Action:  app.CliAction,
		Commands: app.subCommands,
	}

	return cli.Run(os.Args)
}

func (app *App) AddAppCliFlags() []cli.Flag {
	flags := make([]cli.Flag, 0)
	flags = append(flags, &cli.StringFlag{
		Name:  "config,c",
		Value: "config.toml",
		Usage: "config file path",
	})

	for _, flag := range app.extraFlags {
		flags = append(flags, flag)
	}

	return flags
}
