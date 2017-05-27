package handlers

import (
	"log"
	"os"

	"github.com/docker/docker/client"
	"github.com/play-with-docker/play-with-docker/docker"
	"github.com/play-with-docker/play-with-docker/pwd"
)

var core pwd.PWDApi
var Broadcast pwd.BroadcastApi

func Bootstrap() {
	c, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	d := docker.NewDocker(c)

	Broadcast, err = pwd.NewBroadcast(WS, WSError)
	if err != nil {
		log.Fatal(err)
	}

	t := pwd.NewScheduler(Broadcast, d)

	s := pwd.NewStorage()

	core = pwd.NewPWD(d, t, Broadcast, s)

	err = core.SessionLoadAndPrepare()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error decoding sessions from disk ", err)
	}
}
