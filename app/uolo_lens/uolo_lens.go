package main

import (
"artemis/uolo_lens"
"artemis/base/uolo"
)

func main() {

	notifier := uolo_lens.NewUoloLens()

	app := uolo.NewApp(notifier)
	app.WithHttpServer(notifier).
		Run()
}

