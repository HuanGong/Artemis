package main

import (
	"artemis/uolo_center"
	"artemis/base/uolo"
)


func main() {

	notifier := uolo_center.NewNotifier()

	app := uolo.NewApp(notifier)
	app.WithHttpServer(notifier).
		Run()
}
