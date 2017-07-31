package handlers

import (
	"log"
	"os"

	"github.com/docker/docker/client"
	"github.com/play-with-docker/play-with-docker/config"
	"github.com/play-with-docker/play-with-docker/docker"
	"github.com/play-with-docker/play-with-docker/pwd"
	"github.com/play-with-docker/play-with-docker/storage"
)

var core pwd.PWDApi
var broadcast pwd.BroadcastApi

func Bootstrap() {
	c, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	d := docker.NewDocker(c)

	broadcast, err = pwd.NewBroadcast(WS, WSError)
	if err != nil {
		log.Fatal(err)
	}

	t := pwd.NewScheduler(broadcast, d)

	s, err := storage.NewFileStorage(config.SessionsFile)

	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error decoding sessions from disk ", err)
	}
	core = pwd.NewPWD(d, t, broadcast, s)

}
