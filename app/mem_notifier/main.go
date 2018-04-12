package main

import (
	"artemis/mem_notifier"
	"artemis/base/uolo"
)


func main() {

	notifier := mem_notifier.NewNotifier()

	app := uolo.NewApp(notifier)
	app.WithHttpServer(notifier).
		Run()
}
