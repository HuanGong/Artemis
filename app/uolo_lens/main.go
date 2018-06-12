package main

import (
	"artemis/base/uolo"
	"artemis/uolo_lens"
	"github.com/sirupsen/logrus"
)

func main() {

	notifier := uolo_lens.NewUoloLens()

	app := uolo.NewApp(notifier)

	app.WithHttpServer(notifier).
		WithPrometheusMetrics().
		Run()
	logrus.Debugln("App Lens End")
}
