package main

import (
	"artemis/base/uolo"
	"artemis/statisticians"
)


func main() {

	impl := statisticians.NewStatistician()

	app := uolo.NewApp(impl)
	app.WithHttpServer(impl).
		Run()
}
