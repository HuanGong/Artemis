package main

import (
	"artemis/base/uolo"
	"artemis/echo_face"
)

func main() {

	echoface := echo_face.NewEchoFace()

	app := uolo.NewApp(echoface)
	app.WithHttpServer(echoface).
		Run()
}
