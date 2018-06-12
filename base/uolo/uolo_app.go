package uolo

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
	"net"
	"net/http"
	"os"
	"sync"
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
		OnHttpServerInitialized(echo *echo.Echo) error
	}

	StartFlow struct {
		Name    string
		RunType FlowRunType
		Start   func() error
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

func (app *App) WithPrometheusMetrics() *App {
	flow := &StartFlow{
		Name:    "Prometheus Metrics",
		RunType: kRunTypeAsync,
		Start: func() error {
			listener, err := net.Listen("tcp", "0.0.0.0:3009")
			if err != nil {
				return errors.Wrapf(err, "RunPrometheusMatrix listen failed")
			}

			//grpc_prometheus.EnableHandlingTimeHistogram()

			log.Info("=============================================")
			log.Infof("=== Prometheus Metrics start on %s ===", listener.Addr())
			log.Info("=============================================")

			http.Handle("/metrics", promhttp.Handler())
			panic(http.Serve(listener, nil))

			return nil
		},
	}

	app.startFlow = append(app.startFlow, flow)
	return app
}

func (app *App) WithHttpServer(content HttpServerContent) *App {
	flow := &StartFlow{
		Name:    "HttpServer",
		RunType: kRunTypeAsync,
		Start: func() error {
			e := echo.New()
			e.HideBanner = true

			e.HTTPErrorHandler = HttpErrorHandler

			if err := content.OnHttpServerInitialized(e); err != nil {
				log.Errorf("Start HttpServer %s Failed With Error %s\n", content.HttpServerName(), err.Error())
				return errors.Wrap(err, "Start HttpServer Failed")
			}
			if err := e.Start(content.Endpoint()); err != nil {
				log.Infoln("http server end with error:", err.Error())
			}

			return nil
		},
	}

	app.startFlow = append(app.startFlow, flow)
	return app
}

func (app *App) WithCommonFlow(fun func() error, flowRunType FlowRunType) {
	flow := &StartFlow{
		RunType: flowRunType,
		Start:   fun,
	}
	app.startFlow = append(app.startFlow, flow)
}

func (app *App) CliAction(ctx *cli.Context) error {

	log.Infoln("Uolo Application Start Run CliAction")

	if app.content != nil {
		app.content.OnCliApplicationRun()
	}

	log.Infoln("start run start flows:", len(app.startFlow))

	wg := sync.WaitGroup{}
	wg.Add(1) //avoid empty wg.wait
	for _, flow := range app.startFlow {
		log.Infoln("Run Flow ", flow.Name)
		switch {
		case flow.RunType == kRunTypeAsync:
			wg.Add(1)
			go func(startFlow *StartFlow) {
				if err := startFlow.Start(); err != nil {
					log.Panicln("Start sync Flow Failed", err.Error())
				}
				wg.Done()
			}(flow)
		case flow.RunType == kRunTypeSync:
			if err := flow.Start(); err != nil {
				log.Panicln("Start async Flow Failed", err.Error())
			}
		}
	}
	wg.Done()
	wg.Wait()

	return nil
}

func (app *App) Run() error {
	cli := &cli.App{
		Name:     app.name,
		Version:  app.version,
		Usage:    app.description,
		Flags:    app.AddAppCliFlags(),
		Action:   app.CliAction,
		Commands: app.subCommands,
	}

	app.content.BeforeCliRun()
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

func HttpErrorHandler(err error, c echo.Context) {
	log.Debugln(fmt.Sprintf("Request %s from ip %s occur error %s", c.Path(), c.RealIP(), err.Error()))

	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		if he.Inner != nil {
			msg = fmt.Sprintf("%v, %v", err, he.Inner)
		}
	} else {
		msg = http.StatusText(code)
	}
	if _, ok := msg.(string); ok {
		msg = map[string]interface{}{"message": msg}
	}

	if !c.Response().Committed {
		if c.Request().Method == "HEAD" { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
	}
}
