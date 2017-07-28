package task

import (
	"context"
	"fmt"
	"log"

	"github.com/play-with-docker/play-with-docker/docker"
	"github.com/play-with-docker/play-with-docker/event"
	"github.com/play-with-docker/play-with-docker/pwd/types"
)

type DockerSwarmPorts struct {
	Manager   string   `json:"manager"`
	Instances []string `json:"instances"`
	Ports     []int    `json:"ports"`
}

type checkSwarmPorts struct {
	event   event.EventApi
	factory docker.FactoryApi
}

var CheckSwarmPortsEvent event.EventType

func init() {
	CheckSwarmPortsEvent = event.NewEventType("instance docker swarm ports")
}

func (t *checkSwarmPorts) Name() string {
	return "CheckSwarmPorts"
}

func (t *checkSwarmPorts) Run(ctx context.Context, instance *types.Instance) error {
	dockerClient, err := t.factory.GetForInstance(instance.SessionId, instance.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	status, err := getDockerSwarmStatus(ctx, dockerClient)
	if err != nil {
		log.Println(err)
		return err
	}

	if !status.IsManager {
		return nil
	}

	hosts, ps, err := dockerClient.GetSwarmPorts()
	if err != nil {
		log.Println(err)
		return err
	}
	instances := make([]string, len(hosts))
	sessionPrefix := instance.SessionId[:8]
	for i, host := range hosts {
		instances[i] = fmt.Sprintf("%s_%s", sessionPrefix, host)
	}
	ports := make([]int, len(ps))
	for i, port := range ps {
		ports[i] = int(port)
	}

	t.event.Emit(CheckSwarmPortsEvent, instance.SessionId, DockerSwarmPorts{Manager: instance.Name, Instances: instances, Ports: ports})
	return nil
}

func NewCheckSwarmPorts(e event.EventApi, f docker.FactoryApi) *checkSwarmPorts {
	return &checkSwarmPorts{event: e, factory: f}
}
