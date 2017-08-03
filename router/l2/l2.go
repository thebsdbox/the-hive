package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/play-with-docker/play-with-docker/config"
	"github.com/play-with-docker/play-with-docker/router"
)

func director(protocol router.Protocol, host string) (*net.TCPAddr, error) {
	info, err := router.DecodeHost(host)
	if err != nil {
		return nil, err
	}

	port := info.Port

	if info.EncodedPort > 0 {
		port = info.EncodedPort
	}

	if port == 0 {
		if protocol == router.ProtocolHTTP {
			port = 80
		} else if protocol == router.ProtocolHTTPS {
			port = 443
		} else if protocol == router.ProtocolSSH {
			port = 22
		} else if protocol == router.ProtocolDNS {
			port = 53
		}
	}

	t, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", info.InstanceIP, port))
	if err != nil {
		return nil, err
	}
	return t, nil
}

func connectNetworks() error {
	ctx := context.Background()
	c, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	f, err := os.Open(config.SessionsFile)
	if err != nil {
		return err
	}
	defer f.Close()

	networks := map[string]*network.EndpointSettings{}

	err = json.NewDecoder(f).Decode(&networks)
	if err != nil {
		return err
	}

	for netId, opts := range networks {
		settings := &network.EndpointSettings{}
		settings.IPAddress = opts.IPAddress
		log.Printf("Connected to network [%s] with ip [%s]\n", netId, opts.IPAddress)
		c.NetworkConnect(ctx, netId, config.PWDContainerName, settings)
	}

	return nil
}

func monitorNetworks() {
	c, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	ctx := context.Background()

	args := filters.NewArgs()

	cmsg, _ := c.Events(ctx, types.EventsOptions{Filters: args})
	for {
		select {
		case m := <-cmsg:
			if m.Type == "network" {
				// Router has been connected to a new network. Let's get all connections and store them in case of restart.
				container, err := c.ContainerInspect(ctx, config.PWDContainerName)
				if err != nil {
					log.Println(err)
					return
				}

				f, err := os.Create(config.SessionsFile)
				if err != nil {
					log.Println(err)
					return
				}
				err = json.NewEncoder(f).Encode(container.NetworkSettings.Networks)
				if err != nil {
					log.Println(err)
					return
				}
				log.Println("Saved networks")
			}
		}
	}
}

func main() {
	config.ParseFlags()

	err := connectNetworks()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("connect networks:", err)
	}
	go monitorNetworks()

	r := router.NewRouter(director, config.SSHKeyPath)
	r.ListenAndWait(":443", ":53", ":22")
	defer r.Close()
}
