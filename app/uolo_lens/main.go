package main

import (
	"artemis/base/uolo"
	"artemis/uolo_lens"
)

func main() {

	notifier := uolo_lens.NewUoloLens()

	app := uolo.NewApp(notifier)
	app.WithHttpServer(notifier).
		Run()
}
