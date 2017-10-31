package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/docker/docker/client"
	"github.com/play-with-docker/play-with-docker/config"
	"github.com/shirou/gopsutil/load"
)

func Ping(rw http.ResponseWriter, req *http.Request) {
	// Get system load average of the last 5 minutes and compare it against a threashold.

	c, err := client.NewEnvClient()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := c.Info(ctx); err != nil && err == context.DeadlineExceeded {
		log.Printf("Docker info took to long to respond\n")
		rw.WriteHeader(http.StatusGatewayTimeout)
	}

	a, err := load.Avg()
	if err != nil {
		log.Println("Cannot get system load average!", err)
	} else {
		if a.Load5 > config.MaxLoadAvg {
			log.Printf("System load average is too high [%f]\n", a.Load5)
			rw.WriteHeader(http.StatusInsufficientStorage)
		}
	}
}
