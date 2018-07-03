package main

import (
	"artemis/uolo_center"
	"artemis/base/uolo"
)

func main() {

	auCenter := uolo_center.NewUoloCenter()

	app := uolo.NewApp(auCenter)
	app.WithHttpServer(auCenter).
		Run()
}
