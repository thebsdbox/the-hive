package main

import (
	"log"
	"os"
	"time"

	"github.com/thebsdbox/the-hive/config"
	"github.com/thebsdbox/the-hive/docker"
	"github.com/thebsdbox/the-hive/event"
	"github.com/thebsdbox/the-hive/handlers"
	"github.com/thebsdbox/the-hive/id"
	"github.com/thebsdbox/the-hive/k8s"
	"github.com/thebsdbox/the-hive/provisioner"
	"github.com/thebsdbox/the-hive/pwd"
	"github.com/thebsdbox/the-hive/pwd/types"
	"github.com/thebsdbox/the-hive/scheduler"
	"github.com/thebsdbox/the-hive/scheduler/task"
	"github.com/thebsdbox/the-hive/storage"
)

func main() {
	config.ParseFlags()

	e := initEvent()
	s := initStorage()
	df := initDockerFactory(s)
	kf := initK8sFactory(s)

	ipf := provisioner.NewInstanceProvisionerFactory(provisioner.NewWindowsASG(df, s), provisioner.NewDinD(id.XIDGenerator{}, df, s))
	sp := provisioner.NewOverlaySessionProvisioner(df)

	core := pwd.NewPWD(df, e, s, sp, ipf)

	tasks := []scheduler.Task{
		task.NewCheckPorts(e, df),
		task.NewCheckSwarmPorts(e, df),
		task.NewCheckSwarmStatus(e, df),
		task.NewCollectStats(e, df, s),
		task.NewCheckK8sClusterStatus(e, kf),
		task.NewCheckK8sClusterExposedPorts(e, kf),
	}
	sch, err := scheduler.NewScheduler(tasks, s, e, core)
	if err != nil {
		log.Fatal("Error initializing the scheduler: ", err)
	}

	sch.Start()

	d, err := time.ParseDuration(config.PlaygroundLifetime)
	if err != nil {
		log.Fatalf("Cannot parse duration Got: %v", err)
	}

	playground := types.Playground{Domain: config.PlaygroundDomain, DefaultDinDInstanceImage: config.DefaultImage, AvailableDinDInstanceImages: []string{"thebsdbox/dind:game", "thebsdbox/dind:ebpf", "thebsdbox/dind"}, AllowWindowsInstances: config.NoWindows, DefaultSessionDuration: d, Extras: map[string]interface{}{"LoginRedirect": "http://localhost:3000"}, Privileged: true, MaxInstances: config.MaxInstances, AssetsDir: config.UX}
	if _, err := core.PlaygroundNew(playground); err != nil {
		log.Fatalf("Cannot create default playground. Got: %v", err)
	}

	handlers.Bootstrap(core, e)
	handlers.Register(nil)
}

func initStorage() storage.StorageApi {
	s, err := storage.NewFileStorage(config.SessionsFile)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error initializing StorageAPI: ", err)
	}
	return s
}

func initEvent() event.EventApi {
	return event.NewLocalBroker()
}

func initDockerFactory(s storage.StorageApi) docker.FactoryApi {
	return docker.NewLocalCachedFactory(s)
}

func initK8sFactory(s storage.StorageApi) k8s.FactoryApi {
	return k8s.NewLocalCachedFactory(s)
}
